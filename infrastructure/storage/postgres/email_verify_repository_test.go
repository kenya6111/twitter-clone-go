package postgres

import (
	"context"
	"testing"
	"time"
	"twitter-clone-go/domain"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestEmailVerifyRepository_FindByToken(t *testing.T) {
	tests := []struct {
		name         string
		token        string
		want         *domain.EmailVerifyToken
		setupContext func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context
		wantErr      bool
	}{
		{
			name:  "success",
			token: "abc123xyztoken",
			want: &domain.EmailVerifyToken{
				ID:        "1",
				UserID:    "1",
				Token:     "abc123xyztoken",
				ExpiresAt: time.Now(),
				CreatedAt: time.Now(),
			},
			setupContext: func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context {
				return loadWithTx(t, context.Background(), db, tx, "./testdata/email_verify_repository/default.sql")
			},
			wantErr: false,
		},
		{
			name:  "error when email_verify_token not found by token",
			token: "not_found_by_token",
			want:  nil,
			setupContext: func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context {
				return loadWithTx(t, context.Background(), db, tx, "./testdata/email_verify_repository/default.sql")
			},
			wantErr: true,
		},
		{
			name:  "error when context is canceled",
			token: "not_found_by_token",
			want:  nil,
			setupContext: func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return WithTx(ctx, tx)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, tx, err := setup(t)
			if err != nil {
				t.Fatalf("failed to setup database: %v", err)
			}
			t.Cleanup(func() {
				if err := cleanup(context.Background(), t, db, tx); err != nil {
					t.Fatalf("failed to cleanup database: %v", err)
				}
			})

			ctx := tt.setupContext(t, db, tx)
			r := NewEmailVerifyRepository(db)
			got, err := r.FindByToken(ctx, tt.token)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if got != nil {
					t.Fatalf("expected nil result on error, got: %+v", got)
				}
				return
			}
			opts := []cmp.Option{
				cmpopts.IgnoreFields(domain.EmailVerifyToken{}, "ID", "CreatedAt", "ExpiresAt"),
			}

			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Fatalf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
