package router

import "github.com/gorilla/mux"

//Generate returns the application routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
