package postgres

import (
	"context"
	"fmt"
	"testing"
	"twitter-clone-go/domain"

	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestUserRepository_FindByEmail(t *testing.T) {
	ps, err := domain.NewPassword("hashed_PW1!")
	if err != nil {
		panic(1)
	}
	tests := []struct {
		name         string
		email        string
		want         *domain.User
		setupContext func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context
		wantErr      bool
	}{
		{
			name:  "success",
			email: "user1@example.com",
			want: &domain.User{
				ID:       "1",
				Name:     "user1",
				Email:    "user1@example.com",
				Password: ps,
				IsActive: domain.UserStatusActive,
			},
			setupContext: func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context {
				return loadWithTx(t, context.Background(), db, tx, "./testdata/user_repository/default.sql")
			},
			wantErr: false,
		},
		{
			name:  "error when user not found by email",
			email: "notfound@example.com",
			want:  nil,
			setupContext: func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context {
				return loadWithTx(t, context.Background(), db, tx, "./testdata/user_repository/default.sql")
			},
			wantErr: true,
		},
		{
			name:  "error when context is canceled",
			email: "user1@example.com",
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

			r := NewUserRepository(db)
			ctx := tt.setupContext(t, db, tx)
			got, err := r.FindByEmail(ctx, tt.email)
			if err != nil {
				fmt.Println(err.Error())
			}
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
				cmp.Comparer(func(a, b domain.Password) bool {
					return a.Value() == b.Value()
				}),
			}

			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Fatalf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
