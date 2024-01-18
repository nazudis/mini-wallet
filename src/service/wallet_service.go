package service

type WalletService interface {
	InitAccount(customerXid string) (token string, err error)
}
