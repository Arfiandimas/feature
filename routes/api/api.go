package api

import (
	"mandiri_payment_gateway_service/base"
	middlewares "mandiri_payment_gateway_service/guard"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct{}

func (a api) Do(router *base.Router) {
	api := router.Group("/api", middlewares.AppMiddleware{}.Do())

	// Just for example
	api.GET("/hello_world", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{
			Status:  true,
			Message: "api hello world!",
		})
	})

}

func Init() api {
	return api{}
}
