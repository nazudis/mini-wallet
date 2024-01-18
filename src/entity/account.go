package entity

import "github.com/google/uuid"

type Account struct {
	BaseEntityUID
	CustomerXid uuid.UUID `json:"customer_xid" gorm:"column:customer_xid;type:char(36);index"`
	Timestamp
}
