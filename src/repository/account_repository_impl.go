package repository

import (
	"errors"

	"github.com/nazudis/mini-wallet/src/entity"
	"gorm.io/gorm"
)

type AccountRepositoryImpl struct {
	db *gorm.DB
}

// FirstByCustID implements AccountRepository.
func (r AccountRepositoryImpl) FirstByCustID(custId string) (*entity.Account, error) {
	account := new(entity.Account)
	err := r.db.First(&account, "customer_xid = ?", custId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return account, nil
}

// FirstByID implements AccountRepository.
func (r AccountRepositoryImpl) FirstByID(id string) (*entity.Account, error) {
	account := new(entity.Account)
	err := r.db.First(&account, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return account, nil
}

// Create implements AccountRepository.
func (r AccountRepositoryImpl) Create(data *entity.Account) error {
	return r.db.Create(&data).Error
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return AccountRepositoryImpl{
		db: db,
	}
}
