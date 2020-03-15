package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/n1try/kithub2/app/common/errors"
	"github.com/n1try/kithub2/app/config"
	"github.com/n1try/kithub2/app/users"
	"net/http"
)

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Next()

		user, _ := c.Get(config.UserKey)
		if user == nil || !user.(*users.User).IsAdmin() {
			c.AbortWithError(http.StatusUnauthorized, errors.Unauthorized{})
			return
		}
	}
}
