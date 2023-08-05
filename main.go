package main

import (
	"go-pos/config"
	"go-pos/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set up the database connection
	db, err := config.NewDB()
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// Initialize the controllers
	authController := controllers.NewAuthController(db)
	productController := controllers.NewProductController(db)
	transactionController := controllers.NewTransactionController(db)

	// Register the API routes
	api := r.Group("/api")
	{
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
			authRoutes.GET("/validate", authController.ValidateToken)
		}

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
	}

	// Start the server
	r.Run(":8080")
}
