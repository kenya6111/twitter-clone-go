package controllers_test

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"twitter-clone-go/repository/testdata"
	db "twitter-clone-go/tutorial"

	"github.com/gin-gonic/gin"
)

func TestGetUserList(t *testing.T) {
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)

	expected := []db.User{
		testdata.UserTestData[0],
	}
	fmt.Println(expected)

	aCon.GetUserListHandler(c)
	for i, got := range expected {
		t.Run(got.Name, func(t *testing.T) {
			fmt.Println("---------")
			fmt.Println(res.Code)
			fmt.Println(got.ID)
			fmt.Println("---------")
			if res.Code == int(got.ID) {
				t.Errorf("get %v but want %v\n", got.ID, expected[i].ID)
			}

		})
	}
}
