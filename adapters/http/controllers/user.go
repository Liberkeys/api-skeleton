package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Liberkeys/api-skeleton/adapters/http/viewmodels"
	"github.com/Liberkeys/api-skeleton/ports/handlers/user"
)

// UserController ...
type UserController struct {
	queryHandler user.QueryHandler
}

// NewUserController ...
func NewUserController(queryHandler user.QueryHandler) *UserController {
	return &UserController{
		queryHandler: queryHandler,
	}
}

// OnUserList ...
func (controller *UserController) OnUserList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Request received on....") // TODO: into middleware

	fmt.Println("ADAPTER controllers.OnUsersGet") // TODO: into middleware

	// Call port
	users, _ := controller.queryHandler.GetAllUsers() // TODO: manage error

	// View Model
	vm := viewmodels.UserToViewModel(users[0])

	// HTTP Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(vm.Value)) // TODO: use VM and middlware for this and handle error

	fmt.Println("Request terminated on....") // TODO: into middleware
}
