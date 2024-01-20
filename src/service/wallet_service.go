package service

import (
	"github.com/nazudis/mini-wallet/src/entity"
	"github.com/shopspring/decimal"
)

type WalletService interface {
	InitAccount(customerXid string) (token string, err error)
	EnableWalletAccount(customerXid string) (*entity.Wallet, error)
	DisableWalletAccount(customerXid string, isDisabled bool) (*entity.Wallet, error)
	GetEnabledWallet(customerXid string) (*entity.Wallet, error)
	Deposit(data TransactionParams) (*entity.Transaction, error)
	Withdraw(data TransactionParams) (*entity.Transaction, error)
	GetTransactions(customerXid string) ([]entity.Transaction, error)
}

type TransactionParams struct {
	CustomerXid string
	Amount      decimal.Decimal
	ReferenceId string
}
