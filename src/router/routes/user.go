package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Controller:             controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Controller:             controllers.ListAllUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/user/{id}",
		Method:                 http.MethodGet,
		Controller:             controllers.ListUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/user/{id}",
		Method:                 http.MethodPut,
		Controller:             controllers.UpdateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/user/{id}",
		Method:                 http.MethodDelete,
		Controller:             controllers.DeleteUser,
		AuthenticationRequired: false,
	},
}
