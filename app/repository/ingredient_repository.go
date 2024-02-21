package repository

import (
	log "github.com/sirupsen/logrus"
	"velio-admin-backend/app/domain/dao"
	"velio-admin-backend/db"
)

type IngredientRepository interface {
	GetAllIngredients() ([]dao.Ingredient, error)
	GetByID(id int) (dao.Ingredient, error)
	Create(ingredient dao.Ingredient) (dao.Ingredient, error)
	Update(ingredient dao.Ingredient) (dao.Ingredient, error)
	Delete(id int) error
}

type IngredientRepositoryImpl struct {
	db *db.PrismaClient
}

func (receiver IngredientRepositoryImpl) GetAllIngredients() ([]db.IngredientModel, error) {
	ctx := config.PClient.Context
	ingredients, err := receiver.db.Ingredient.FindMany().Exec(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return ingredients, nil
}

func IngredientRepositoryInit(db *db.PrismaClient) *IngredientRepositoryImpl {
	return &IngredientRepositoryImpl{
		db: db,
	}
}
