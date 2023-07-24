package user_service

import (
	"mandiri_payment_gateway_service/cmd/repositories/Users"
	"mandiri_payment_gateway_service/internal/response"
)

type FindUserService struct {
	Id int
}

func (s FindUserService) Do() response.ServiceResponse {
	findUser, errorFetch := user_repositories.UserRepository{}.FindUser(s.Id)
	if errorFetch != nil {
		return response.Service().Error("error", errorFetch.Error())
	}
	return response.Service().Success("success", findUser)
}
