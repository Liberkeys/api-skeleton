package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Liberkeys/api-skeleton/adapters/http/viewmodels"
	"github.com/Liberkeys/api-skeleton/infrastructure/application"
	"github.com/Liberkeys/api-skeleton/ports/handlers/user"
)

// OnUserList ...
func OnUserList(ctx *application.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Request received on....") // TODO: into middleware

	fmt.Println("ADAPTER controllers.OnUsersGet") // TODO: into middleware

	// Call port
	users, _ := user.GetAllUsers(ctx) // TODO: manage error

	// View Model
	vm := viewmodels.ToUsersViewModel(users)
	buffer, _ := json.Marshal(vm) // TODO handle error

	// HTTP Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	_, _ = w.Write(buffer) // TODO: use VM and middlware for this and handle error

	fmt.Println("Request terminated on....") // TODO: into middleware
}

// OnUserFetch ...
func OnUserFetch(ctx *application.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, _ := user.GetOneUser(ctx, "b91a3b87-8545-466d-8a8e-aeeb7a9270e3") // TODO: manage error and fetch id from URL
	vm := viewmodels.ToUserViewModel(user)

	buffer, _ := json.Marshal(vm) // TODO handle error

	// HTTP Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	_, _ = w.Write(buffer) // TODO: use VM and middlware for this and handle error

	fmt.Println("Request terminated on....") // TODO: into middleware
}

// OnUserCreate ...
func OnUserCreate(ctx *application.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// OnUserUpdate ...
func OnUserUpdate(ctx *application.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// OnUserDelete ...
func OnUserDelete(ctx *application.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
