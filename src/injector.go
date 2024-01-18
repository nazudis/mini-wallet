//go:build wireinject
// +build wireinject

package src

import (
	"github.com/google/wire"
	"github.com/nazudis/mini-wallet/engine/restapi/controller"
	"github.com/nazudis/mini-wallet/src/database"
	"github.com/nazudis/mini-wallet/src/repository"
	"github.com/nazudis/mini-wallet/src/service"
)

func InitializeWalletController() (controller.WalletController, error) {
	wire.Build(
		controller.NewWalletController,
		service.NewWalletService,
		repository.NewAccountRepository,
		database.GetDB,
	)

	return nil, nil
}
