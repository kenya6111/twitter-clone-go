package http

import (
	"errors"
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
	case apperrors.Unauthorized:
		statusCode = http.StatusUnauthorized
	case apperrors.NoTargetData, apperrors.ReqBodyDecodeFailed, apperrors.BadParam, apperrors.DuplicateData:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	c.JSON(statusCode, appErr)
}
