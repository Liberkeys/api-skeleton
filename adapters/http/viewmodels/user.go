package viewmodels

import (
	"github.com/Liberkeys/api-skeleton/ports/models"
)

// UserViewModel ...
type UserViewModel struct {
	Value string `json:"value"`
}

// UsersViewModel ...
type UsersViewModel struct {
	Users []UserViewModel `json:"users"`
}

// ToUsersViewModel ...
func ToUsersViewModel(users []*models.User) UsersViewModel {
	vm := UsersViewModel{}
	for i := range users {
		vm.Users = append(vm.Users, ToUserViewModel(users[i]))
	}
	return vm
}

// ToUserViewModel ...
func ToUserViewModel(user *models.User) UserViewModel {
	return UserViewModel{
		Value: user.Name,
	}
}
