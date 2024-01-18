package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/nazudis/mini-wallet/src/helper"
	"gorm.io/gorm"
)

type Account struct {
	BaseEntityUID
	CustomerXid uuid.UUID `json:"customer_xid" gorm:"column:customer_xid;type:char(36);index;unique"`
	Timestamp
}

func (m *Account) AfterCreate(db *gorm.DB) error {
	wallet := &Wallet{
		OwnedBy:    m.CustomerXid,
		Status:     WalletStatusDisabled,
		DisabledAt: helper.VarToPointer(time.Now()),
	}
	return db.Create(&wallet).Error
}
