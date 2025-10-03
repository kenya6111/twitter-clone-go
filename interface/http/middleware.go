package http

import (
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
			c.Abort()
			return
		}
		if len(cmd) == 0 {
			c.Redirect(202, "/")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
