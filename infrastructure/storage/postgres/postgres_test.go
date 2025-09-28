package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"
	"twitter-clone-go/apperrors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BuildDBConfig_Test() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "mypassword",
		Port:     "5432",
		DBName:   "testdb",
	}
	return &dbConfig
}

func DbURL_Test(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)
}

func SetupDB_Test() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), DbURL_Test(BuildDBConfig_Test()))
	return db, err
}

func setup(t *testing.T) (db *pgxpool.Pool, tx pgx.Tx, err error) { // OK
	db, err = SetupDB_Test()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	tx, err = db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}
	return db, tx, nil
}

func cleanup(ctx context.Context, t *testing.T, db *pgxpool.Pool, tx pgx.Tx) error {
	// if err := tx; err != nil {
	// 	t.Errorf("transaction error: %v", err)
	// }
	if err := tx.Rollback(ctx); err != nil {
		t.Fatalf("failed to rollback transaction: %v", err)
	}
	db.Close()
	return nil
}

// テストごとに欲しいDBデータを生成する。あとcontextも返す
func loadWithTx(t *testing.T, ctx context.Context, db *pgxpool.Pool, tx pgx.Tx, paths ...string) context.Context {
	ctx = WithTx(ctx, tx)

	// パスが指定されていない場合はデフォルトのパスを使用
	if len(paths) == 0 {
		paths = []string{"./testdata/default.sql"}
	}

	// 複数のSQLファイルを順次実行
	for _, path := range paths {
		if err := loadFixtures(ctx, tx, db, path, t); err != nil {
			t.Fatalf("failed to load fixtures from %s: %v", path, err)
		}
	}

	defer func() {
		if err := recover(); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				panic(err)
			}
			panic(err)
		}
	}()
	return ctx
}

func loadFixtures(ctx context.Context, tx pgx.Tx, db *pgxpool.Pool, path string, t *testing.T) error {
	sql, err := os.ReadFile(path)

	if err != nil {
		return apperrors.NAData.Wrap(err, "fail to insert user ")
	}
	if _, err = tx.Exec(ctx, string(sql)); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("failed to exec fixture from %s: %v", path, err)
		return err
	}

	return nil
}
