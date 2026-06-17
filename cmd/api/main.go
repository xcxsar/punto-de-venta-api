package main

import (
	"log"
	"net/http"
	"os"
	"pos-api/internal/database"
	"pos-api/internal/handlers"
	"pos-api/internal/repository"

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

	productRepo := repository.NewProductRepository(db)

	productHandler := handlers.NewProductHandler(productRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.HandleFunc("GET /products", productHandler.GetProducts)
	mux.HandleFunc("GET /products/{id}", productHandler.GetProductByID)
	mux.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
