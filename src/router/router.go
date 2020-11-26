package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Generate returns the application routes
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configurate(r)
}
