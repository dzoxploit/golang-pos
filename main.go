package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"go-pos/config"
	"go-pos/controllers"
	"go-pos/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func generateRandomKey() (string, error) {
	key := make([]byte, 32) // 32 bytes for a secure key
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

func main() {
	secretKey, err := generateRandomKey()
	if err != nil {
		fmt.Println("Failed to generate JWT secret key:", err)
		return
	}

	// Set the secret key as an environment variable
	os.Setenv("JWT_SECRET_KEY", secretKey)
	
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()

	db, err := config.NewDB()
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// Initialize the controllers
	authController := controllers.NewAuthController(db)
	productController := controllers.NewProductController(db)
	transactionController := controllers.NewTransactionController(db)
	personController := controllers.NewPersonController(db)


	
	// Register the API routes

	authRoutes := r.Group("/auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
			authRoutes.GET("/validate", authController.ValidateToken)
		}

	api := r.Group("/api")
	api.Use(middlewares.Authenticate())
	{

		productRoutes := api.Group("/products")
		{
			productRoutes.GET("/", productController.ListProducts)
			productRoutes.POST("/create", productController.CreateProduct)
			productRoutes.GET("/:id", productController.GetProduct)
			productRoutes.PUT("/:id", productController.UpdateProduct)
			productRoutes.DELETE("/:id", productController.DeleteProduct)
		}

		transactionRoutes := api.Group("/transactions")
		{
			transactionRoutes.POST("/", transactionController.CreateTransaction)
			transactionRoutes.GET("/", transactionController.ListTransactions)
		}

		personRoutes := api.Group("/persons")
		{
			personRoutes.GET("/", personController.ListPersons)
			personRoutes.POST("/create", personController.CreatePerson)
			personRoutes.GET("/:id", personController.GetPerson)
			personRoutes.PUT("/:id", personController.UpdatePerson)
			personRoutes.DELETE("/:id", personController.DeletePerson)
		}
	}

	// Start the server
	r.Run(":7000")
}
