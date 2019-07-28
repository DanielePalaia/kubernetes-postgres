package main

// Main routes for get, post, put and delete
import (
	"kubernetes-postgres/engine"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	// Returns all elements of the collection todos
	Route{
		"getCollection",
		"GET",
		"/todos",
		engine.GetCollection,
	},
	// Get an element of a collection
	Route{
		"getElement",
		"GET",
		"/todos/{todoId}",
		engine.GetElement,
	},
	// Create element in a collection
	Route{
		"createElement",
		"PUT",
		"/todos",
		engine.CreateElement,
	},
	// Update element of a collection
	Route{
		"updateElement",
		"POST",
		"/todos",
		engine.UpdateElement,
	},
	// Delete all elements of the collection
	Route{
		"deleteCollection",
		"DELETE",
		"/todos",
		engine.DeleteCollection,
	},
	// Delete an element of a collection
	Route{
		"deleteElement",
		"DELETE",
		"/todos/{todoId}",
		engine.DeleteElement,
	},
	// Delete an element of a collection
	/*Route{
		"swagger",
		"GET",
		"/swagger",
		httpSwagger.WrapHandler,
	},*/
}
