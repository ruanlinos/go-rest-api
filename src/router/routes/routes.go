package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents all api routes.
type Route struct {
	URI                    string
	Method                 string
	Controller             func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

//Configurate returns all routes configurated.
func Configurate(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Controller).Methods(route.Method)
	}

	return r
}
