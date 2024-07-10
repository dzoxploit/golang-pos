package repositories

import (
	"go-pos/models"

	"github.com/jinzhu/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) CreatePerson(product *models.Person) (*models.Person, error) {
	// Create the product in the database
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *PersonRepository) GetPersonByID(id uint) (*models.Person, error) {
	// Get the product from the database based on the ID
	var person models.Person
	if err := r.db.Where("id = ?", id).First(&person).Error; err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *PersonRepository) UpdatePerson(person *models.Person) error {
	// Update the product in the database
	if err := r.db.Save(person).Error; err != nil {
		return err
	}

	return nil
}


func (r *PersonRepository) DeletePerson(id uint) error {
	// Delete the product from the database based on the ID
	if err := r.db.Where("id = ?", id).Delete(models.Person{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) ListPerson() ([]models.Person, error) {
	// Get all products from the database
	var persons []models.Person
	if err := r.db.Find(&persons).Error; err != nil {
		return nil, err
	}

	return persons, nil
}
