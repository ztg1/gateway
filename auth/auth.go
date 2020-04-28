package auth

import (
	"errors"
	"gateway/srv/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, msg.NewMessageVo(201, errors.New("token not found")))
			return
		}

		c.Next()
	}
}
