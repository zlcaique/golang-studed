package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Gen vai retonar um router com as rotas configuradas
func Gen() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
