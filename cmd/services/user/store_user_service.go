package user_service

import (
	user_repositories "mandiri_payment_gateway_service/cmd/repositories/user"
	user_entities "mandiri_payment_gateway_service/cmd/repositories/user/entities"
	user_request "mandiri_payment_gateway_service/cmd/requests/user"
	"mandiri_payment_gateway_service/internal/response"
)

type StoreUserService struct {
	Request user_request.UserStoreRequest
}

func (s StoreUserService) Do() response.ServiceResponse {
	addUser, errorAdd := user_repositories.UserRepository{}.StoreUser(user_entities.Users{
		Name:       s.Request.Name,
		Email:      s.Request.Email,
		Alamat:     s.Request.Alamat,
		Usia:       s.Request.Usia,
		Pendidikan: s.Request.Pendidikan,
	})
	if errorAdd != nil {
		return response.Service().Error("error", errorAdd.Error())
	}
	return response.Service().Success("success", addUser)
}
