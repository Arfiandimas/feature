package user_controller

import (
	"github.com/gin-gonic/gin"
	user_service "mandiri_payment_gateway_service/cmd/services/user"
	"mandiri_payment_gateway_service/internal/response"
	"strconv"
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

func (c UserController) Find(ctx *gin.Context) {
	id := ctx.Param("id")
	intVal, _ := strconv.Atoi(id)
	service := user_service.FindUserService{
		Id: intVal,
	}.Do()
	if !service.Status {
		ctx.JSON(400, response.Api().Error(service.Message, "Get", service.Data))
		return
	}
	ctx.JSON(400, response.Api().Success(service.Message, "Get", service.Data))
	return
}
