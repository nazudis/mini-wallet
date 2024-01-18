package controller

import "net/http"

type WalletController interface {
	InitAccount(w http.ResponseWriter, r *http.Request)
}
