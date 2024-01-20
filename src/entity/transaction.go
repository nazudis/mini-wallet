package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	TrxStatusSuccess = "success"
	TrxStatusProcess = "process"
	TrxStatusFailed  = "failed"

	TrxDeposit    = "deposit"
	TrxWithdrawal = "withdrawal"
)

type Transaction struct {
	BaseEntityUID
	OwnedBy      uuid.UUID       `json:"owned_by" gorm:"column:owned_by;type:char(36);index"`
	ReferenceId  uuid.UUID       `json:"reference_id" gorm:"column:reference_id;type:char(36);index"`
	Amount       decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(64,15);default:0"`
	Status       string          `json:"status" gorm:"column:status"`
	TransactedAt time.Time       `json:"transacted_at" gorm:"column:transacted_at;type:timestamp;not null"`
	Type         string          `json:"type" gorm:"column:type"`
	Timestamp
}
