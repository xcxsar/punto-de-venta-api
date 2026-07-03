package main

import (
	"log"
	"net/http"
	"os"
	"pos-api/internal/database"
	"pos-api/internal/handlers"
	"pos-api/internal/models"
	"pos-api/internal/repositories"

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

	productHandler := handlers.NewProductHandler(productRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.HandleFunc("GET /products", productHandler.GetProducts)
	mux.HandleFunc("GET /products/{id}", productHandler.GetProductByID)
	mux.HandleFunc("GET /products/category/{category_id}", productHandler.GetProductsByCategory)
	mux.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)

	mux.HandleFunc("POST /categories", categoryHandler.CreateCategory)
	mux.HandleFunc("GET /categories", categoryHandler.GetCategories)
	mux.HandleFunc("PUT /categories/{id}", categoryHandler.UpdateCategory)
	mux.HandleFunc("DELETE /categories/{id}", categoryHandler.DeleteCategory)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
