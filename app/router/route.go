package router

import (
	"github.com/gin-gonic/gin"
	"velio-admin-backend/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		ingredient := api.Group("/ingredient")
		ingredient.GET("", init.IngredientCtrl.GetAllIngredients)
		//user.POST("", init.UserCtrl.AddUserData)
		//user.GET("/:userID", init.UserCtrl.GetUserById)
		//user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		//user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	}

	return router
}
