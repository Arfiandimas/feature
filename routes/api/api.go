package api

import (
	"mandiri_payment_gateway_service/base"
	user_controller "mandiri_payment_gateway_service/cmd/controllers/user"
	middlewares "mandiri_payment_gateway_service/guard"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func (a Api) Do(router *base.Router) {
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

	apiUser := api.Group("/user")
	apiUser.GET("/", user_controller.UserController{}.Get)

}

func Init() Api {
	return Api{}
}
