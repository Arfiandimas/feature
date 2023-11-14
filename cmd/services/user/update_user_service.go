package user_service

import (
	user_repositories "mandiri_payment_gateway_service/cmd/repositories/user"
	user_entities "mandiri_payment_gateway_service/cmd/repositories/user/entities"
	user_request "mandiri_payment_gateway_service/cmd/requests/user"
	"mandiri_payment_gateway_service/internal/response"
)

type UpdateUserService struct {
	Request user_request.UserUpdateRequest
	Id      int
}

func (s UpdateUserService) Do() response.ServiceResponse {
	updateUser, errorUpdate := user_repositories.UserRepository{}.UpdateUser(user_entities.Users{
		Name:       s.Request.Name,
		Email:      s.Request.Email,
		Alamat:     s.Request.Alamat,
		Usia:       s.Request.Usia,
		Pendidikan: s.Request.Pendidikan,
	}, s.Id)
	if errorUpdate != nil {
		return response.Service().Error("error", errorUpdate.Error())
	}
	return response.Service().Success("success", updateUser)
}
