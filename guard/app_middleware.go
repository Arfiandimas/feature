package middlewares

import (
	"mandiri_payment_gateway_service/internal/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AppMiddleware struct {
}

func (m AppMiddleware) Do() gin.HandlerFunc {

	return func(c *gin.Context) {

		appKey := c.GetHeader("app-key")

		if os.Getenv("APP_KEY") != appKey || appKey == "" {

			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Api().Error("you can't do this", "mandiri_middleware", nil))
			return

		}

	}
}
