package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
	"github.com/ishu17077/hectoclash-backend/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.POST("/signup", controllers.SignUp())
	incomingRoutes.POST("/login", controllers.Login())
	incomingRoutes.Use(middlewares.Authenticate())

	incomingRoutes.GET("/user", controllers.GetUser())

}
