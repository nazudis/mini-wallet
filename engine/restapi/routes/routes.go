package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nazudis/mini-wallet/src/helper"
)

func AppRoutes(r chi.Router) {

	// Base endpoint
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res := helper.PlugResponse(w)
		res.ReplySuccess(map[string]any{"name": "mini-wallet", "service": "ok"})
	})
}