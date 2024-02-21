package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"velio-admin-backend/app/constant"
	"velio-admin-backend/app/pkg"
	"velio-admin-backend/app/repository"
)

type IngredientService interface {
	GetAllIngredients() (c *gin.Context)
	GetByID() (c *gin.Context)
	Create() (c *gin.Context)
	Update() (c *gin.Context)
	Delete() (c *gin.Context)
}

type IngredientServiceImpl struct {
	ingredientRepository repository.IngredientRepository
}

func (receiver IngredientServiceImpl) GetAllIngredients(c *gin.Context) {
	defer pkg.PanicHandler(c)
	log.Info("Start to execute program for get all ingredient")

	data, err := receiver.ingredientRepository.GetAllIngredients()

	if err != nil {
		log.Error("Error while getting all ingredient", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func IngredientServiceInit(ingredientRepository repository.IngredientRepository) *IngredientServiceImpl {
	return &IngredientServiceImpl{
		ingredientRepository: ingredientRepository,
	}
}
