package services

import "go-pos/models"

type MockPersonService struct {
	// Add any necessary fields for your mock implementation
}

func (s *MockPersonService) CreatePerson(person *models.Person) (*models.Person, error) {
	// Mock implementation for CreatePerson...
	return &models.Person{
		ID:   1,
		Name: "Test Person",
		Age:  30,
	}, nil
}

func (s *MockPersonService) GetPersonByID(id string) (*models.Person, error) {
	// Mock implementation for GetPersonByID...
	return &models.Person{
		ID:   1,
		Name: "Test Person",
		Age:  30,
	}, nil
}

func (s *MockPersonService) UpdatePerson(person *models.Person) error {
	// Mock implementation for UpdatePerson...
	return nil
}

func (s *MockPersonService) DeletePerson(id string) error {
	// Mock implementation for DeletePerson...
	return nil
}

func (s *MockPersonService) ListPersons() ([]models.Person, error) {
	// Mock implementation for ListPersons...
	return []models.Person{
		{
			ID:   1,
			Name: "Test Person 1",
			Age:  30,
		},
		{
			ID:   2,
			Name: "Test Person 2",
			Age:  25,
		},
	}, nil
}
