package repository_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"twitter-clone-go/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var repo *repository.MyAppRepository

// 全テスト共通の前処理を書く
func setup() error {
	var err error
	pool, err = pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1",
		"5432",
		"postgres",
		"mypassword",
		"testdb",
	))
	if err != nil {
		log.Fatal(err)
	}
	repo = repository.NewMyAppRepository(pool)
	return nil
}

func teardown() {
	pool.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
	teardown()
}
