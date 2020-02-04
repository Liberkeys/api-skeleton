package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Liberkeys/api-skeleton/infrastructure/app"
	"github.com/julienschmidt/httprouter"
)

// HTTPList ...
func HTTPList(ctx *app.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Request received on....")

	fmt.Println("ADAPTER controllers.OnUsersGet")

	// Call port
	users, _ := GetAllUsers(ctx)

	// View Model
	vm := toUsersViewModel(users)
	buffer, _ := json.Marshal(vm)

	// HTTP Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	_, _ = w.Write(buffer) // TODO: use VM and middlware for this and handle error

	fmt.Println("Request terminated on....")
}

// HTTPFetch ...
func HTTPFetch(ctx *app.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, _ := GetOneUser(ctx, "b91a3b87-8545-466d-8a8e-aeeb7a9270e3") // TODO: manage error and fetch id from URL
	vm := toUserViewModel(user)

	buffer, _ := json.Marshal(vm)

	// HTTP Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	_, _ = w.Write(buffer) // TODO: use VM and middlware for this and handle error

	fmt.Println("Request terminated on....")
}

// HTTPCreate ...
func HTTPCreate(ctx *app.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// HTTPUpdate ...
func HTTPUpdate(ctx *app.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// HTTPDelete ...
func HTTPDelete(ctx *app.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
