package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezkyal/simple-go-login/entity/config"
	"github.com/rezkyal/simple-go-login/utils"
)

func JwtAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c, cfg.Token.Secret)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
