package controller

import (
	"github.com/gin-gonic/gin"
)

type IngredientController interface {
	GetAllIngredients(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IngredientControllerImpl struct {
	svc service.IngredientService
}

func (receiver IngredientControllerImpl) GetAllIngredients(c *gin.Context) {
	receiver.svc.GetAllIngredients()
}

func IngredientControllerInit(ingredientService service.IngredientService) *IngredientControllerImpl {
	return &IngredientControllerImpl{
		svc: ingredientService,
	}
}
