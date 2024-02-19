package main

import (
	"apiv1/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.BookRoutes(router)

	http.ListenAndServe(":8080", router)
}
