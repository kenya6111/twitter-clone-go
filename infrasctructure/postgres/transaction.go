package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Transaction はトランザクション管理を提供する
type Transaction struct {
	client *client
}

// NewTransaction は新しいトランザクションインスタンスを作成する
func NewTransaction(pool *pgxpool.Pool) *Transaction {
	return &Transaction{client: &client{pool: pool}}
}

// Do はトランザクション内で関数を実行する
func (t *Transaction) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.client.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// context に tx を埋め込んで渡す
	if err := fn(WithTx(ctx, tx)); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// DoWithLock は排他制御付きでトランザクション内で関数を実行する
// func (t *Transaction) DoWithLock(ctx context.Context, fn func(ctx context.Context) error) error {
// 	return t.client.pool.BeginFunc(ctx, func(tx pgx.Tx) error {
// 		// 必要に応じて排他制御のロジックを追加
// 		return fn(WithTx(ctx, tx))
// 	})
// }
