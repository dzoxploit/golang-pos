package controllers

import (
	"go-pos/models"
	"go-pos/services"
	"go-pos/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TransactionController struct {
	transactionService *services.TransactionService
	productService *services.ProductService
}

func NewTransactionController(db *gorm.DB) *TransactionController {
	return &TransactionController{
		transactionService: services.NewTransactionService(db),
	}
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
    var transaction models.Transaction
    if err := ctx.ShouldBindJSON(&transaction); err != nil {
        utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Create the transaction
    createdTransaction, err := c.transactionService.CreateTransaction(&transaction)
    if err != nil {
        utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create transaction")
        return
    }

    // Subtract the stock quantity of the products involved in the transaction
    err = c.productService.UpdateProductStock(createdTransaction)
    if err != nil {
        utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update product stock")
        return
    }

    utils.SendSuccessResponse(ctx, http.StatusCreated, "Transaction created successfully", createdTransaction)
}

func (c *TransactionController) ListTransactions(ctx *gin.Context) {
	transactions, err := c.transactionService.ListTransactions()
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch transactions")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusOK, "Transactions fetched successfully", transactions)
}
