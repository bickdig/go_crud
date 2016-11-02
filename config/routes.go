// routes.go
// Copyright (C) 2016 Reza Jatnika <rezajatnika@gmail.com>
//
// Distributed under terms of the MIT license.
//

package config

import (
	"net/http"
	"os"

	"github.com/bickdig/go_crud/app/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Route define the data structure of a Route
type Route struct {
	Name    string
	Methods string
	Pattern string
	Handler http.HandlerFunc
}

// Routes define the data structure of Routes
type Routes []Route

var (
	postController = controllers.NewPostController()
	routes         = []Route{
		Route{
			"GET /",
			"GET",
			"/",
			postController.Index,
		},
		Route{
			"GET /posts/new",
			"GET",
			"/posts/new",
			postController.New,
		},
		Route{
			"GET /posts/{id}",
			"GET",
			"/posts/{id}",
			postController.Show,
		},
		Route{
			"POST /posts",
			"POST",
			"/posts",
			postController.Create,
		},
		Route{
			"GET /posts/{id}/edit",
			"GET",
			"/posts/{id}/edit",
			postController.Edit,
		},
		Route{
			"PUT /posts/{id}",
			"POST",
			"/posts/{id}",
			postController.Update,
		},
	}
)

// NewRouter returns a new mux.Router struct
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Methods).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	return router
}

func NewLoggedRouter() http.Handler {
	loggedRouter := handlers.LoggingHandler(os.Stdout, NewRouter())
	return loggedRouter
}
