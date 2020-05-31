package route

import (
	"github.com/vpatel95/go-api-server/app/api"

	"github.com/gorilla/mux"
)

var (
	Router *mux.Router
)

func SetRoutes() {
	Router = mux.NewRouter()
	api.SetRoutes(Router)
}
