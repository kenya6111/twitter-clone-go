package postgres

import (
	"context"
	"fmt"
	"os"

	db "twitter-clone-go/tutorial"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

type DBConfig struct {
	Host     string
	User     string
	DBName   string
	Password string
	Port     string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)
}

func SetupDB() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), DbURL(BuildDBConfig()))
	return db, err
}

// txKey はcontextにトランザクションを保存するためのキー
type txKey struct{}

// WithTx はcontextにトランザクションを設定する
func WithTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// GetTx はcontextからトランザクションを取得する
func GetTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(pgx.Tx)
	return tx, ok
}

// client はデータベースクライアントを表す。client という言葉は「サービスを利用する側」「依頼する側」という意味
// つまり Go アプリから見れば、Postgres に問い合わせをする「クライアント」が必要なので、その役割を持った構造体に client という名前を付けている
type client struct {
	pool *pgxpool.Pool
}

// DB はcontextからトランザクションを取得し、なければプールを返す
func (c *client) DB(ctx context.Context) interface{} {
	if tx, ok := GetTx(ctx); ok {
		return tx
	}
	return c.pool
}

// Querier はcontextに応じて適切なquerierを返す
func (c *client) Querier(ctx context.Context) *db.Queries {
	tx := c.DB(ctx)
	if tx, ok := tx.(pgx.Tx); ok {
		return db.New(tx)
	}
	return db.New(c.pool)
}
