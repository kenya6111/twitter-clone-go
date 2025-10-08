package http

import (
	"net/http"
	"twitter-clone-go/infrastructure/session_store"

	"github.com/gin-gonic/gin"
)

const (
	cookieKey = "sid"
)

func CheckLogin(s *session_store.SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		redisKey, _ := c.Cookie(cookieKey)
		cmd, err := s.Client.Get(c.Request.Context(), "session:"+redisKey).Result()
		if err != nil {

			c.AbortWithStatusJSON(http.StatusUnauthorized, APIResponse{
				ErrCode: "S013",
				Message: "session expired or not found",
			})
			return
		}
		if len(cmd) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, APIResponse{
				ErrCode: "S013",
				Message: "session expired or not found",
			})
			return
		} else {
			c.Next()
		}
	}
}
