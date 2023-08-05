package test

import (
	"bytes"
	"encoding/json"
	"go-pos/controllers"
	"go-pos/models"
	"go-pos/services"
	"go-pos/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProductController_CreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()

	// Create a new mock product service
	mockProductService := &services.MockProductService{}
	productController := controllers.NewProductController(mockProductService)

	// Define the route
	router.POST("/products/create", productController.CreateProduct)

	// Define the test case
	t.Run("Create product with valid data", func(t *testing.T) {
		// Create a request
		productData := &models.Product{
			Name:     "Test Product",
			Price:    12.34,
			Quantity: 100,
		}
		reqJSON, _ := json.Marshal(productData)
		req, _ := http.NewRequest(http.MethodPost, "/products/create", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Check the response
		assert.Equal(t, http.StatusCreated, w.Code)

		var response utils.SuccessResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, "Product created successfully", response.Message)
		// Additional checks for the response data if needed
	})
	// Add more test cases for different scenarios
}
