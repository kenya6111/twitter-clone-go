package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	ErrCode string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		ErrCode: "0",
		Message: "success",
		Data:    data,
	})
}
