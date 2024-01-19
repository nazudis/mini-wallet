package service

import "github.com/nazudis/mini-wallet/src/entity"

type WalletService interface {
	InitAccount(customerXid string) (token string, err error)
	EnableWalletAccount(customerXid string) (*entity.Wallet, error)
	GetEnabledWallet(customerXid string) (*entity.Wallet, error)
}
