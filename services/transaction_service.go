package services

import (
	"go-pos/models"
	"go-pos/repositories"

	"github.com/jinzhu/gorm"
)

type TransactionService struct {
	transactionRepository *repositories.TransactionRepository
	productRepository *repositories.ProductRepository
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		transactionRepository: repositories.NewTransactionRepository(db),
	}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return s.transactionRepository.CreateTransaction(transaction)
}

func (s *TransactionService) UpdateProductStock(transaction *models.Transaction) error {
    for _, item := range transaction.Items {
        // Get the product by its ID
        product, err := s.productRepository.GetProductByID(item.ProductID)
        if err != nil {
            return err
        }

        // Subtract the quantity from the product's stock
        product.Quantity -= item.Quantity

        // Update the product's stock in the repository
        err = s.productRepository.UpdateProduct(product)
        if err != nil {
            return err
        }
    }

    return nil
}

func (s *TransactionService) ListTransactions() ([]*models.Transaction, error) {
	transactions, err := s.transactionRepository.ListTransactions()
	if err != nil {
		return nil, err
	}

	// Convert []models.Transaction to []*models.Transaction
	var transactionPointers []*models.Transaction
	for i := range transactions {
		transactionPointers = append(transactionPointers, &transactions[i])
	}

	return transactionPointers, nil
}
