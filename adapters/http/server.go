package http

import (
	"fmt"
	"net/http"

	"github.com/Liberkeys/api-skeleton/domains/users"
	"github.com/Liberkeys/api-skeleton/infrastructure/app"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	router http.Handler
}

// NewServer ...
func NewServer(ctx *app.Context) *Server {
	router := httprouter.New()

	router.GET("/users", wrapEndpoint(ctx, users.HTTPList))
	router.POST("/users", wrapEndpoint(ctx, users.HTTPCreate))
	router.GET("/users/:id", wrapEndpoint(ctx, users.HTTPFetch))
	router.POST("/users/:id", wrapEndpoint(ctx, users.HTTPUpdate))
	router.DELETE("/users/:id", wrapEndpoint(ctx, users.HTTPDelete))

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
	ctx *app.Context,
	endpoint func(ctx *app.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params),
) func(http.ResponseWriter, *http.Request, httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// TODO Generate Request ID and Trace ID
		endpoint(ctx, w, r, p)
	}
}
