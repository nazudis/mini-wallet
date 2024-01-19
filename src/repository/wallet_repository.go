package repository

import "github.com/nazudis/mini-wallet/src/entity"

type WalletRepository interface {
	FirstByCustID(custId string) (*entity.Wallet, error)
	EnabledWallet(data *entity.Wallet) error
	DisabledWallet(data *entity.Wallet) error
}
