package user_controller

import (
	"github.com/gin-gonic/gin"
	user_service "mandiri_payment_gateway_service/cmd/services/user"
	"mandiri_payment_gateway_service/internal/response"
)

type UserController struct {
}

func (c UserController) Get(ctx *gin.Context) {

	service := user_service.GetUserService{}.Do()
	if !service.Status {
		ctx.JSON(400, response.Api().Error(service.Message, "Get", service.Data))
		return
	}
	ctx.JSON(400, response.Api().Success(service.Message, "Get", service.Data))
	return
}