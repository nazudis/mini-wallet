package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nazudis/mini-wallet/src/entity"
	"github.com/nazudis/mini-wallet/src/helper"
	"github.com/nazudis/mini-wallet/src/repository"
)

type WalletServiceImpl struct {
	AccountRepository repository.AccountRepository
	WalletRepository  repository.WalletRepository
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

// EnableWalletAccount implements WalletService.
func (s WalletServiceImpl) EnableWalletAccount(customerXid string) (*entity.Wallet, error) {
	_, err := uuid.Parse(customerXid)
	if err != nil {
		return nil, err
	}

	account, err := s.AccountRepository.FirstByCustID(customerXid)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("account not found")
	}

	wallet, err := s.WalletRepository.FirstByCustID(customerXid)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, fmt.Errorf("wallet not found")
	}

	if wallet.Status == entity.WalletStatusEnabled {
		return nil, fmt.Errorf("already enabled")
	}

	err = s.WalletRepository.EnabledWallet(wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// DisableWalletAccount implements WalletService.
func (s WalletServiceImpl) DisableWalletAccount(customerXid string, isDisabled bool) (*entity.Wallet, error) {
	_, err := uuid.Parse(customerXid)
	if err != nil {
		return nil, err
	}

	account, err := s.AccountRepository.FirstByCustID(customerXid)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("account not found")
	}

	wallet, err := s.WalletRepository.FirstByCustID(customerXid)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, fmt.Errorf("wallet not found")
	}

	if isDisabled {
		if wallet.Status == entity.WalletStatusDisabled {
			return nil, fmt.Errorf("already disabled")
		}

		err = s.WalletRepository.DisabledWallet(wallet)
		if err != nil {
			return nil, err
		}
	}

	return wallet, nil
}

// GetEnabledWallet implements WalletService.
func (s WalletServiceImpl) GetEnabledWallet(customerXid string) (*entity.Wallet, error) {
	_, err := uuid.Parse(customerXid)
	if err != nil {
		return nil, err
	}

	wallet, err := s.WalletRepository.FirstByCustID(customerXid)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, fmt.Errorf("wallet not found")
	}

	if wallet.Status == entity.WalletStatusDisabled {
		return nil, fmt.Errorf("wallet disabled")
	}

	return wallet, nil
}

func NewWalletService(accountRepository repository.AccountRepository, walletRepository repository.WalletRepository) WalletService {
	return WalletServiceImpl{
		AccountRepository: accountRepository,
		WalletRepository:  walletRepository,
	}
}
