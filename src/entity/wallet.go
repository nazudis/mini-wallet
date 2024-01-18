package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	WalletStatusEnabled  = "enabled"
	WalletStatusDisabled = "disabled"
)

type Wallet struct {
	BaseEntityUID
	OwnedBy    uuid.UUID       `json:"owned_by" gorm:"column:owned_by;type:char(36);index"`
	Balance    decimal.Decimal `json:"balance" gorm:"column:balance;type:decimal(64,15);default:0"`
	Status     string          `json:"status" gorm:"column:status;default:disabled"`
	EnabledAt  *time.Time      `json:"enabled_at" gorm:"column:enabled_at;type:timestamp"`
	DisabledAt *time.Time      `json:"disabled_at" gorm:"column:disabled_at;type:timestamp"`
	Timestamp
}
