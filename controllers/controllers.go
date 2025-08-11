package controllers

import "twitter-clone-go/controllers/services"

type MyAppController struct {
	svc services.SessionServicer
}

func NewMyAppController(s services.SessionServicer) *MyAppController {
	return &MyAppController{svc: s}
}
