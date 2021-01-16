package infrastructure

import (
	"github.com/gin-gonic/gin"

	"api/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	// User
	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("api/v1/users", func(c *gin.Context) { userController.Create(c) })

	Router = router
}
