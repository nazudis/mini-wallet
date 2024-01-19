package repository

import (
	"errors"
	"time"

	"github.com/nazudis/mini-wallet/src/entity"
	"github.com/nazudis/mini-wallet/src/helper"
	"gorm.io/gorm"
)

type WalletRepositoryImpl struct {
	db *gorm.DB
}

// FirstByCustID implements WalletRepository.
func (r WalletRepositoryImpl) FirstByCustID(custId string) (*entity.Wallet, error) {
	wallet := new(entity.Wallet)
	err := r.db.First(&wallet, "owned_by = ?", custId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return wallet, nil
}

// DisabledWallet implements WalletRepository.
func (r WalletRepositoryImpl) DisabledWallet(data *entity.Wallet) error {
	data.Status = entity.WalletStatusDisabled
	data.DisabledAt = helper.VarToPointer(time.Now())
	data.EnabledAt = nil

	return r.db.Save(&data).Error
}

// EnabledWallet implements WalletRepository.
func (r WalletRepositoryImpl) EnabledWallet(data *entity.Wallet) error {
	data.Status = entity.WalletStatusEnabled
	data.EnabledAt = helper.VarToPointer(time.Now())
	data.DisabledAt = nil

	return r.db.Save(&data).Error
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return WalletRepositoryImpl{
		db: db,
	}
}
