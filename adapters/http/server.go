package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Liberkeys/api-skeleton/adapters/http/controllers"
	"github.com/Liberkeys/api-skeleton/infrastructure/application"
)

type Server struct {
	router http.Handler
}

func NewServer(ctx *application.Context) *Server {
	router := httprouter.New()

	router.GET("/users", wrapEndpoint(ctx, controllers.OnUserList))
	router.POST("/users", wrapEndpoint(ctx, controllers.OnUserCreate))
	router.GET("/users/:id", wrapEndpoint(ctx, controllers.OnUserFetch))
	router.POST("/users/:id", wrapEndpoint(ctx, controllers.OnUserUpdate))
	router.DELETE("/users/:id", wrapEndpoint(ctx, controllers.OnUserDelete))

	return &Server{
		router: router,
	}
}

// Start ...
func (server *Server) Start() error {
	fmt.Println("Start server on :8080")
	return http.ListenAndServe(":8080", server.router)
}

func wrapEndpoint(
	ctx *application.Context,
	endpoint func(ctx *application.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params),
) func(http.ResponseWriter, *http.Request, httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// TODO Generate Request ID and Trace ID
		endpoint(ctx, w, r, p)
	}
}
