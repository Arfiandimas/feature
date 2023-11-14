package user_repositories

import (
	"fmt"
	"mandiri_payment_gateway_service/base"
	"mandiri_payment_gateway_service/cmd/repositories/user/entities"
)

type UserRepository struct {
}

func (r UserRepository) GetUser() ([]user_entities.Users, error) {
	var user []user_entities.Users
	result := base.OpenDB().Gorm().Find(&user)
	return user, result.Error
}

func (r UserRepository) FindUser(id int) (user_entities.Users, error) {
	var user user_entities.Users
	result := base.OpenDB().Gorm().Where("id = ?", id).First(&user)
	return user, result.Error
}

func (r UserRepository) StoreUser(user user_entities.Users) (user_entities.Users, error) {
	result := base.OpenDB().Gorm().Save(&user)
	return user, result.Error
}

func (r UserRepository) UpdateUser(user user_entities.Users, id int) (user_entities.Users, error) {
	fmt.Println(&user, id)
	result := base.OpenDB().Gorm().Where("id = ?", id).Updates(&user)
	return user, result.Error
}
