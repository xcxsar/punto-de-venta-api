package models

import "github.com/shopspring/decimal"

type SaleItem struct {
	ID        uint64          `json:"id" gorm:"primaryKey"`
	ProductID *uint64         `json:"product_id"`
	Quantity  int             `json:"quantity"`
	UnitPrice decimal.Decimal `json:"unit_price"`
	Subtotal  decimal.Decimal `json:"subtotal"`

	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
