package http

import (
	"context"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/application"
	"twitter-clone-go/infrastructure/session_store"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase application.UserUsecase
}

func NewUserHandler(u application.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) Home(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello World!"})
}

func (h *UserHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func (h *UserHandler) GetUserList(c *gin.Context) {
	users, err := h.usecase.GetUserList()
	if err != nil {
		ErrorHandler(c, err)
		return
	}
	SuccessResponse(c, users)
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var request application.SignUpInfo
	if err := c.BindJSON(&request); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		ErrorHandler(c, err)
		return
	}
	if err := h.usecase.SignUp(c.Request.Context(), request); err != nil {
		ErrorHandler(c, err)
		return
	}
	SuccessResponse(c, nil)
}
func (h *UserHandler) Activate(c *gin.Context) {
	token := c.Query("token")
	if err := h.usecase.Activate(c.Request.Context(), token); err != nil {
		ErrorHandler(c, err)
		return
	}
	SuccessResponse(c, nil)
}

func (h *UserHandler) Login(c *gin.Context) {
	var request application.LoginInfo
	if err := c.BindJSON(&request); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		ErrorHandler(c, err)
		return
	}
	ctx := context.WithValue(c.Request.Context(), session_store.GinContextKey, c)

	user, err := h.usecase.Login(ctx, request)
	if err != nil {
		// emailの存在チェックをされないように、一律のエラーを返す
		ErrorHandler(c, apperrors.ErrUnauthorized)
		return
	}
	SuccessResponse(c, user)
}

func (h *UserHandler) Logout(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), session_store.GinContextKey, c)
	err := h.usecase.Logout(ctx)
	if err != nil {
		// emailの存在チェックをされないように、一律のエラーを返す
		ErrorHandler(c, apperrors.ErrLogout)
		return
	}
	SuccessResponse(c, nil)

}
