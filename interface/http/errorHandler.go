package http

import (
	"errors"
	"log"
	"net/http"
	"twitter-clone-go/apperrors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error) {
	var appErr *apperrors.MyAppError
	if !errors.As(err, &appErr) {
		appErr = &apperrors.MyAppError{
			ErrCode: apperrors.Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int

	switch appErr.ErrCode {
	case apperrors.NAData:
		statusCode = http.StatusNotFound
	case apperrors.AuthUnauthorized:
		statusCode = http.StatusUnauthorized
	case apperrors.NoTargetData, apperrors.ReqBodyDecodeFailed, apperrors.ReqBadParam, apperrors.DuplicateData:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	log.Println(err.Error())
	c.JSON(statusCode, appErr)
}
