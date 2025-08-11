package repository_test

import (
	"fmt"
	"testing"
	"twitter-clone-go/repository/testdata"
	db "twitter-clone-go/tutorial"
)

func TestGetUserList(t *testing.T) {

	expected := []db.User{
		testdata.UserTestData[0],
		testdata.UserTestData[1],
		testdata.UserTestData[2],
		testdata.UserTestData[3],
	}

	gots, err := repo.GetUserList()
	if err != nil {
		t.Fatal(err)
	}
	for i, got := range gots {
		t.Run(got.Name, func(t *testing.T) {
			fmt.Println("!", i)
			fmt.Println(got.ID)
			fmt.Println(expected[i].ID)

			if got.ID != expected[i].ID {
				t.Errorf("get %v but want %v\n", got.ID, expected[i].ID)
			}

		})
	}
}
