package postgres

import (
	"context"
	"testing"
	"twitter-clone-go/domain"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

			ctx := tt.setupContext(t, db, tx)
			r := NewUserRepository(db)
			got, err := r.FindByEmail(ctx, tt.email)
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

func TestUserRepository_CountByEmail(t *testing.T) {
	tests := []struct {
		name         string
		email        string
		want         int64
		setupContext func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context
		wantErr      bool
	}{
		{
			name:  "success when user not registered",
			email: "user11111@example.com",
			want:  0,
			setupContext: func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context {
				return loadWithTx(t, context.Background(), db, tx, "./testdata/user_repository/default.sql")
			},
			wantErr: false,
		},
		{
			name:  "error when user already resisted (count > 0) ",
			email: "user1@example.com",
			want:  1,
			setupContext: func(t *testing.T, db *pgxpool.Pool, tx pgx.Tx) context.Context {
				return loadWithTx(t, context.Background(), db, tx, "./testdata/user_repository/default.sql")
			},
			wantErr: false,
		},
		{
			name:  "error when context is canceled",
			email: "user1@example.com",
			want:  1,
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
			got, err := r.CountByEmail(ctx, tt.email)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error but got none")
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
func TestUserRepository_CreateUser(t *testing.T) {
	ps, err := domain.NewPassword("hashed_PW1!")
	if err != nil {
		panic(1)
	}
	tests := []struct {
		name    string
		user    *domain.User
		want    *domain.User
		wantErr bool
	}{
		{
			name: "succes",
			user: &domain.User{
				Name:     "新規ユーザー",
				Password: ps,
				Email:    "unique-newuser@example.com",
				IsActive: domain.UserStatusInactive,
			},
			want: &domain.User{
				Name:     "新規ユーザー",
				Password: ps,
				Email:    "unique-newuser@example.com",
				IsActive: domain.UserStatusInactive,
			},
			wantErr: false,
		},
		{
			name: "error when email is not unique",
			user: &domain.User{
				Name:     "既存ユーザー",
				Password: ps,
				Email:    "user1@example.com",
				IsActive: domain.UserStatusInactive,
			},
			want:    nil,
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
			ctx := loadWithTx(t, context.Background(), db, tx, "./testdata/user_repository/default.sql")
			r := NewUserRepository(db)
			got, err := r.CreateUser(ctx, tt.user.Name, tt.user.Email, tt.user.Password.Value())
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if got != nil {
					t.Fatalf("expected nil result on error, got: %+v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// IDとタイムスタンプを無視して比較
			opts := []cmp.Option{
				cmp.Comparer(func(a, b domain.Password) bool {
					return a.Value() == b.Value()
				}),
				cmpopts.IgnoreFields(domain.User{}, "ID"),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Fatalf("mismatch (-got +want):\n%s", diff)
			}
			// IDが設定されていることを確認
			if got.ID == "" {
				t.Fatal("ID should be set")
			}
		})
	}
}
func TestUserRepository_UpdateUser(t *testing.T) {

	ps, err := domain.NewPassword("hashed_PW1!")
	if err != nil {
		panic(1)
	}
	tests := []struct {
		name    string
		userId  string
		want    *domain.User
		wantErr bool
	}{
		{
			name:   "succes",
			userId: "1",
			want: &domain.User{
				Name:     "user1",
				Password: ps,
				Email:    "user1@example.com",
				IsActive: domain.UserStatusActive,
			},
			wantErr: false,
		},
		{
			name:    "error when userId is not exist",
			userId:  "11111",
			want:    nil,
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
			ctx := loadWithTx(t, context.Background(), db, tx, "./testdata/user_repository/default.sql")
			r := NewUserRepository(db)
			got, err := r.UpdateUser(ctx, tt.userId)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if got != nil {
					t.Fatalf("expected nil result on error, got: %+v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// IDとタイムスタンプを無視して比較
			opts := []cmp.Option{
				cmp.Comparer(func(a, b domain.Password) bool {
					return a.Value() == b.Value()
				}),
				cmpopts.IgnoreFields(domain.User{}, "ID"),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Fatalf("mismatch (-got +want):\n%s", diff)
			}
			// IDが設定されていることを確認
			if got.ID == "" {
				t.Fatal("ID should be set")
			}
		})
	}
}
