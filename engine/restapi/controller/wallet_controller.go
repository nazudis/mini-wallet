package controller

import "net/http"

type WalletController interface {
	InitAccount(w http.ResponseWriter, r *http.Request)
	EnableWalletAccount(w http.ResponseWriter, r *http.Request)
	DisableWalletAccount(w http.ResponseWriter, r *http.Request)
	GetWallet(w http.ResponseWriter, r *http.Request)
	Deposit(w http.ResponseWriter, r *http.Request)
	Withdraw(w http.ResponseWriter, r *http.Request)
	GetWalletTransactions(w http.ResponseWriter, r *http.Request)
}
