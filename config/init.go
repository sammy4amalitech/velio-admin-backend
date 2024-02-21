package config

import (
	"velio-admin-backend/app/controller"
	"velio-admin-backend/app/repository"
	"velio-admin-backend/app/service"
)

type Initialization struct {
	ingredientRepo repository.IngredientRepository
	ingredientSvc  service.IngredientService
	ingredientCtrl controller.IngredientController
}

func NewInitialization(
	ingredientRepo repository.IngredientRepository,
	ingredientCtrl controller.IngredientController,
	ingredientService service.IngredientService,
) *Initialization {
	return &Initialization{
		ingredientRepo: ingredientRepo,
		ingredientSvc:  ingredientService,
		ingredientCtrl: ingredientCtrl,
	}
}
