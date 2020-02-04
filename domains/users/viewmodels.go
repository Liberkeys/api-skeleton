package users

import (
	"github.com/Liberkeys/api-skeleton/models")

// UserViewModel ...
type userViewModel struct {
	Value string `json:"value"`
}

// UsersViewModel ...
type usersViewModel struct {
	Users []userViewModel `json:"users"`
}

// ToUsersViewModel ...
func toUsersViewModel(users []*models.User) usersViewModel {
	vm := usersViewModel{}
	for i := range users {
		vm.Users = append(vm.Users, toUserViewModel(users[i]))
	}
	return vm
}

// ToUserViewModel ...
func toUserViewModel(user *models.User) userViewModel {
	return userViewModel{
		Value: user.Name,
	}
}
