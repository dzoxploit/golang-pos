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

func TestTransactionController_CreateTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()

	// Create a new mock transaction service
	mockTransactionService := &services.MockTransactionService{}
	transactionController := controllers.NewTransactionController(mockTransactionService)

	// Define the route
	router.POST("/transactions/create", transactionController.CreateTransaction)

	// Define the test case
	t.Run("Create transaction with valid data", func(t *testing.T) {
		// Create a request
		transactionData := &models.Transaction{
			// Initialize transaction data with valid values
		}
		reqJSON, _ := json.Marshal(transactionData)
		req, _ := http.NewRequest(http.MethodPost, "/transactions/create", bytes.NewBuffer(reqJSON))
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
		assert.Equal(t, "Transaction created successfully", response.Message)
		// Additional checks for the response data if needed
	})
	// Add more test cases for different scenarios
}
