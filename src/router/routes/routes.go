package routes

import "net/http"

//Route represents all api routes.
type Route struct {
	URI                    string
	Method                 string
	Controller             func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}
