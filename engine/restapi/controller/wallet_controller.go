package controller

import "net/http"

type WalletController interface {
	InitAccount(w http.ResponseWriter, r *http.Request)
	EnableWalletAccount(w http.ResponseWriter, r *http.Request)
	DisableWalletAccount(w http.ResponseWriter, r *http.Request)
	GetWallet(w http.ResponseWriter, r *http.Request)
}
