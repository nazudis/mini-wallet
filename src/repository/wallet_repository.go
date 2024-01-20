package repository

import (
	"github.com/nazudis/mini-wallet/src/entity"
)

type WalletRepository interface {
	FirstByCustID(custId string) (*entity.Wallet, error)
	EnabledWallet(data *entity.Wallet) error
	DisabledWallet(data *entity.Wallet) error

	// Transaction
	Deposit(wallet *entity.Wallet, trx *entity.Transaction) error
	Withdraw(wallet *entity.Wallet, trx *entity.Transaction) error
	GetTransactions(custId string) []entity.Transaction
}
