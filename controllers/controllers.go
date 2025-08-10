package controllers

import "twitter-clone-go/services"

type MyAppController struct {
	svc *services.MyAppService
}

func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{svc: s}
}
