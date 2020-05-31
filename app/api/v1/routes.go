package apiv1

import (
	"github.com/gorilla/mux"

	"github.com/vpatel95/go-api-server/app/api/v1/auth"
	"github.com/vpatel95/go-api-server/app/api/v1/users"
)

func SetRoutes(router *mux.Router) {
	route := router.PathPrefix("/v1").Subrouter()

	auth.Routes(route)
	users.Routes(route)
}
