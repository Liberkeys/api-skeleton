package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Liberkeys/api-skeleton/adapters/http/controllers"
)

type Server struct {
	router http.Handler
}

func NewServer(userContoller *controllers.UserController) *Server {
	router := httprouter.New()
	router.GET("/users", userContoller.OnUserList)

	return &Server{
		router: router,
	}
}

// Start ...
func (server *Server) Start() error {
	fmt.Println("Start server on :8080")
	return http.ListenAndServe(":8080", server.router)
}
