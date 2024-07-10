package services

import (
	"go-pos/models"

	"github.com/jinzhu/gorm"
)

type PersonService struct {
	db *gorm.DB
}

func NewPersonService(db *gorm.DB) *PersonService {
	return &PersonService{
		db: db,
	}
}

func (s *PersonService) CreatePerson(person *models.Person) (*models.Person, error) {
	if err := s.db.Create(person).Error; err != nil {
		return nil, err
	}
	return person, nil
}

func (s *PersonService) GetPersonByID(id string) (*models.Person, error) {
	var person models.Person
	if err := s.db.Where("id = ?", id).First(&person).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

func (s *PersonService) UpdatePerson(person *models.Person) error {
	if err := s.db.Save(person).Error; err != nil {
		return err
	}
	return nil
}

func (s *PersonService) DeletePerson(id string) error {
	if err := s.db.Where("id = ?", id).Delete(&models.Person{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *PersonService) ListPersons() ([]models.Person, error) {
	var persons []models.Person
	if err := s.db.Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}
