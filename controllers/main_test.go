package controllers_test

import (
	"os"
	"testing"
	"twitter-clone-go/controllers"
	"twitter-clone-go/controllers/testdata"
)

var aCon *controllers.MyAppController

func setup() error {
	svc := testdata.NewServiceMock()
	aCon = controllers.NewMyAppController(svc)
	return nil
}

func teardown() {
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
	teardown()
}
