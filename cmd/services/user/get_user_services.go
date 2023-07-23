package user_service

import (
	"mandiri_payment_gateway_service/cmd/repositories/Users"
	"mandiri_payment_gateway_service/internal/response"
)

type GetUserService struct {
}

func (s GetUserService) Do() response.ServiceResponse {
	getUsers, _ := user_repositories.UserRepository{}.GetUser()
	return response.Service().Success("success", getUsers)
}
