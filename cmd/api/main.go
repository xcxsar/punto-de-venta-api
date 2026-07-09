package main

import (
	"log"
	"os"
	"pos-api/internal/database"
	"pos-api/internal/handlers"
	"pos-api/internal/models"
	"pos-api/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found.")
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	log.Println("Successfully connected to the database via GORM.")

	productRepo := repositories.NewBaseRepository[models.Product](db)
	categoryRepo := repositories.NewBaseRepository[models.Category](db)
	saleRepo := repositories.NewBaseRepository[models.Sale](db)

	productHandler := handlers.NewProductHandler(productRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)
	saleHandler := handlers.NewSaleHandler(saleRepo)

	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.GetProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
		api.GET("/products/category/:category_id", productHandler.GetProductsByCategory)

		api.POST("/categories", categoryHandler.CreateCategory)
		api.GET("/categories", categoryHandler.GetCategories)
		api.PUT("/categories/:id", categoryHandler.UpdateCategory)
		api.DELETE("/categories/:id", categoryHandler.DeleteCategory)

		api.POST("/sales", saleHandler.CreateSale)
		api.GET("/sales", saleHandler.GetSales)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	r.Run(":" + port)
}
