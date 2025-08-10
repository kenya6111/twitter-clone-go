package services

import (
	"twitter-clone-go/repository"
)

type MyAppService struct {
	// 2. フィールドに sql.DB 型を含める
	repo *repository.MyAppRepository
}

// コンストラクタの定義
func NewMyAppService(repo *repository.MyAppRepository) *MyAppService {
	return &MyAppService{repo: repo}
}
