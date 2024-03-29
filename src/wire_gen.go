// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package src

import (
	"github.com/nazudis/mini-wallet/engine/restapi/controller"
	"github.com/nazudis/mini-wallet/src/database"
	"github.com/nazudis/mini-wallet/src/repository"
	"github.com/nazudis/mini-wallet/src/service"
)

// Injectors from injector.go:

func InitializeWalletController() (controller.WalletController, error) {
	db := database.GetDB()
	accountRepository := repository.NewAccountRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	walletService := service.NewWalletService(accountRepository, walletRepository)
	walletController := controller.NewWalletController(walletService)
	return walletController, nil
}
