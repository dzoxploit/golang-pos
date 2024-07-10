package controllers

import (
	"go-pos/models"
	"go-pos/services"
	"go-pos/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PersonController struct {
	personService *services.PersonService
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		personService: services.NewPersonService(db),
	}
}

func NewPersonControllerTest(personService *services.PersonService) *PersonController {
	return &PersonController{
		personService: personService,
	}
}

func (c *PersonController) CreatePerson(ctx *gin.Context) {
	var person models.Person
	if err := ctx.ShouldBindJSON(&person); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newPerson, err := c.personService.CreatePerson(&person)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusCreated, "Person created successfully", newPerson)
}

func (c *PersonController) GetPerson(ctx *gin.Context) {
	personID := ctx.Param("id")
	person, err := c.personService.GetPersonByID(personID)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusNotFound, "Person Not Found")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusOK, "Person found successfully", person)
}

func (c *PersonController) UpdatePerson(ctx *gin.Context) {
	personID := ctx.Param("id")
	var updatedPerson models.Person
	if err := ctx.ShouldBindJSON(&updatedPerson); err != nil {
		utils.SendErrorResponse(ctx, 400, "Invalid request payload")
		return
	}

	// Convert the personID (string) to a uint
	parsedPersonID, err := strconv.ParseUint(personID, 10, 64)
	if err != nil {
		utils.SendErrorResponse(ctx, 400, "Invalid person ID")
		return
	}

	// Assign the personID to the ID field of updatedPerson
	updatedPerson.ID = uint(parsedPersonID)

	// Call the UpdatePerson function from the service
	err = c.personService.UpdatePerson(&updatedPerson)
	if err != nil {
		utils.SendErrorResponse(ctx, 500, "Failed to update person")
		return
	}

	utils.SendSuccessResponse(ctx, 200, "Person updated successfully", updatedPerson)
}

func (c *PersonController) DeletePerson(ctx *gin.Context) {
	personID := ctx.Param("id")
	err := c.personService.DeletePerson(personID)
	if err != nil {
		utils.SendErrorResponse(ctx, 500, "Failed to delete person")
		return
	}

	utils.SendSuccessResponse(ctx, 200, "Person deleted successfully", nil)
}

func (c *PersonController) ListPersons(ctx *gin.Context) {
	persons, err := c.personService.ListPersons()
	if err != nil {
		utils.SendErrorResponse(ctx, 500, "Failed to fetch persons")
		return
	}

	utils.SendSuccessResponse(ctx, 200, "Persons fetched successfully", persons)
}
