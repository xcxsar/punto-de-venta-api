package handlers

import (
	"fmt"
	"net/http"
	"pos-api/internal/models"
	"pos-api/internal/repositories"

	"github.com/gin-gonic/gin"
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
	PaymentMethod string `json:"payment_method" binding:"required"`
	Items         []struct {
		ProductID uint64 `json:"product_id" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required,min=1"`
	} `json:"items" binding:"required,min=1"`
}

func (h *SaleHandler) CreateSale(c *gin.Context) {
	var req SaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sale)
}

func (h *SaleHandler) GetSales(c *gin.Context) {
	var sales []models.Sale

	if err := h.Repo.DB.Preload("Items").Find(&sales).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sales)
}
