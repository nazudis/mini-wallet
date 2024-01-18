package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nazudis/mini-wallet/engine/restapi/routes"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.DefaultLogger)
	r.Use(middleware.NoCache)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Recoverer)
	r.NotFound(r.NotFoundHandler())

	r.Route("/api/v1", routes.AppRoutes)

	fmt.Println("run on port 8080")
	http.ListenAndServe("localhost:8080", r)
}
