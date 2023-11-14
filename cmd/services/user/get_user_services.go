package user_service

import (
	"mandiri_payment_gateway_service/cmd/repositories/user"
	"mandiri_payment_gateway_service/internal/response"
)

type GetUserService struct {
}

func (s GetUserService) Do() response.ServiceResponse {
	getUsers, errorFetch := user_repositories.UserRepository{}.GetUser()
	if errorFetch != nil {
		return response.Service().Error("error", errorFetch.Error())
	}
	return response.Service().Success("success", getUsers)
}
