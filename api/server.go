// Package api implements rest api for the service
// include http-router, middlwares, validators and controllers
package api

import (
	"net/http"
	"time"

	"github.com/Akshit8/go-boilerplate/api/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type health struct {
	Health string `json:"health"`
	Check  bool   `json:"check"`
}

// CreateNewServer creates a new http router with chi
// configure all middlewares with router
// registers all endpoints to controllers
// returns the router
func CreateNewServer() *chi.Mux {
	r := chi.NewRouter()

	// middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Compress(5, "application/json"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			newhealth := health{
				Health: "this is a health check route",
				Check: true,
			}
			utils.SendHTTPResponse(newhealth, 200, w)
		})
		r.Mount("/users", userRouter())
		r.Mount("/notes", noteRouter())
	})

	return r
}

func userRouter() http.Handler {
	r := chi.NewRouter()
	return r
}

func noteRouter() http.Handler {
	r := chi.NewRouter()
	return r
}
