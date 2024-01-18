package repository

import "github.com/nazudis/mini-wallet/src/entity"

type AccountRepository interface {
	FirstByID(id string) (*entity.Account, error)
	FirstByCustID(custId string) (*entity.Account, error)
	Create(data *entity.Account) error
}
