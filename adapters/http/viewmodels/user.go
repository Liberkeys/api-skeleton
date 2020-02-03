package viewmodels

import (
	"github.com/Liberkeys/api-skeleton/ports/models"
)

// UserViewModel ...
type UserViewModel struct{
	Value string 
}

// UserToViewModel ...
func UserToViewModel(user models.User) UserViewModel {
	vm := UserViewModel{Value: `{"user": "`+ user.Name +`"}`} 
	return vm
}

// UserFromViewModel ...
func UserFromViewModel(raw string) models.User {
	user := models.User{
		Name: raw,
	}
	return user
}