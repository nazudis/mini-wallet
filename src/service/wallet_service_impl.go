package service

import (
	"github.com/google/uuid"
	"github.com/nazudis/mini-wallet/src/entity"
	"github.com/nazudis/mini-wallet/src/helper"
	"github.com/nazudis/mini-wallet/src/repository"
)

type WalletServiceImpl struct {
	AccountRepository repository.AccountRepository
}

// InitAccount implements WalletService.
func (s WalletServiceImpl) InitAccount(customerXid string) (token string, err error) {
	custXid, err := uuid.Parse(customerXid)
	if err != nil {
		return "", err
	}

	account, err := s.AccountRepository.FirstByCustID(customerXid)
	if err != nil {
		return "", err
	}

	if account == nil {
		account = &entity.Account{CustomerXid: custXid}
		err = s.AccountRepository.Create(account)
		if err != nil {
			return "", err
		}
	}

	return helper.GenToken(account.CustomerXid.String()), nil
}

func NewWalletService(accountRepository repository.AccountRepository) WalletService {
	return WalletServiceImpl{
		AccountRepository: accountRepository,
	}
}
