package user_repositories

import (
	"mandiri_payment_gateway_service/base"
	"mandiri_payment_gateway_service/cmd/repositories/Users/entities"
)

type UserRepository struct {
}

func (u UserRepository) GetUser() ([]user_entities.Users, error) {
	var user []user_entities.Users
	result := base.OpenDB().Gorm().Find(&user)
	return user, result.Error
}

func (u UserRepository) FindUser(id int) (user_entities.Users, error) {
	var user user_entities.Users
	result := base.OpenDB().Gorm().Where("id = ?", id).First(&user)
	return user, result.Error
}
