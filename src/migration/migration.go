package migration

import (
	"github.com/nazudis/mini-wallet/src/database"
	"github.com/nazudis/mini-wallet/src/entity"
)

func init() {
	db := database.GetDB()
	err := db.AutoMigrate(entity.Account{}, entity.Wallet{}, entity.Transaction{})
	if err != nil {
		panic(err)
	}
}
