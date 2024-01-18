package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nazudis/mini-wallet/src"
	"github.com/nazudis/mini-wallet/src/helper"
)

func walletRoutes(r chi.Router) {
	walletCtrl, _ := src.InitializeWalletController()

	r.Post("/init", walletCtrl.InitAccount)
	r.Route("/wallet", func(walletR chi.Router) {
		// walletR.Use()
	})
}

func AppRoutes(r chi.Router) {

	// Base endpoint
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res := helper.PlugResponse(w)
		data := helper.MapJSON{"name": "mini-wallet", "service": "ok"}
		res.ReplySuccess(data)
	})

	r.Group(walletRoutes)
}
