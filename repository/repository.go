package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB(p *pgxpool.Pool) {
	pool = p
}

type MyAppRepository struct {
	pool *pgxpool.Pool
}

// コンストラクタの定義
func NewMyAppRepository(pool *pgxpool.Pool) *MyAppRepository {
	return &MyAppRepository{pool: pool}
}
