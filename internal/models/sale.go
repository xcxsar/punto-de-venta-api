package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Sale struct {
	ID            uint64          `json:"id" gorm:"primaryKey"`
	Total         decimal.Decimal `json:"total"`
	PaymentMethod string          `json:"payment_method"`
	CreatedAt     time.Time       `json:"created_at"`

	Items []SaleItem `json:"items,omitempty" gorm:"foreignKey:SaleID"`
}
