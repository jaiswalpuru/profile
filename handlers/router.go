package handlers

import (
	"cmd/models"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(routes *[]models.Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *routes {
		if route.HeaderDefault.HandlerFunc != nil {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HeaderDefault.HandlerFunc).
				Headers(route.HeaderDefault.HeaderType.Name, route.HeaderDefault.HeaderType.Value)
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.DefaultHandler)

	}

	staticDir := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir)))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods(http.MethodGet)
	scriptDir := http.Dir("./scripts")
	scriptFileHandler := http.StripPrefix("/scripts/", http.FileServer(scriptDir))
	router.PathPrefix("/scripts/").Handler(scriptFileHandler).Methods(http.MethodGet)
	return router
}
