package models

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ID       int64           `json:"id" gorm:"primaryKey"`
	Name     string          `json:"name"`
	Stock    int             `json:"stock"`
	Price    decimal.Decimal `json:"price"`
	ImageURL *string         `json:"image_url"`
}

func (Product) TableName() string {
	return "inventory.product"
}
