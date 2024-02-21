// injector.go
package config

import (
	"github.com/google/wire"
	"velio-admin-backend/app/service"
)

var dbSet = wire.NewSet(ConnectDB)

var ingredientRepoSet = wire.NewSet(
	service.IngredientServiceInit,
	wire.Bind(new(service.IngredientService), new(service.IngredientServiceImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, dbSet, ingredientRepoSet)
	return nil
}
