package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pos-api/internal/models"
	"pos-api/internal/repositories"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SaleHandler struct {
	Repo *repositories.BaseRepository[models.Sale]
}

func NewSaleHandler(repo *repositories.BaseRepository[models.Sale]) *SaleHandler {
	return &SaleHandler{Repo: repo}
}

type SaleRequest struct {
	PaymentMethod string `json:"payment_method"`
	Items         []struct {
		ProductID uint64 `json:"product_id"`
		Quantity  int    `json:"quantity"`
	} `json:"items"`
}

func (h *SaleHandler) CreateSale(w http.ResponseWriter, r *http.Request) {
	var req SaleRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(req.Items) == 0 {
		http.Error(w, "No items sold", http.StatusBadRequest)
		return
	}

	sale := models.Sale{
		PaymentMethod: req.PaymentMethod,
		Total:         decimal.Zero,
	}

	err := h.Repo.DB.Transaction(func(tx *gorm.DB) error {
		for _, reqItem := range req.Items {
			var product models.Product

			if err := tx.First(&product, reqItem.ProductID).Error; err != nil {
				return fmt.Errorf("product ID %d not found", reqItem.ProductID)
			}

			if product.Stock == nil || *product.Stock < reqItem.Quantity {
				return fmt.Errorf("not enough stock for %s", product.Name)
			}

			qtyDecimal := decimal.NewFromInt(int64(reqItem.Quantity))
			subTotal := product.Price.Mul(qtyDecimal)

			sale.Total = sale.Total.Add(subTotal)

			productID := reqItem.ProductID

			saleItem := models.SaleItem{
				ProductID: &productID,
				Quantity:  reqItem.Quantity,
				UnitPrice: product.Price,
				Subtotal:  subTotal,
			}

			sale.Items = append(sale.Items, saleItem)

			if err := tx.Model(&models.Product{}).
				Where("id = ?", product.ID).
				Update("stock", gorm.Expr("stock - ?", reqItem.Quantity)).Error; err != nil {
				return err
			}
		}

		if err := tx.Create(&sale).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sale)
}

func (h *SaleHandler) GetSales(w http.ResponseWriter, r *http.Request) {
	var sales []models.Sale

	if err := h.Repo.DB.Preload("Items").Find(&sales).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sales)
}
