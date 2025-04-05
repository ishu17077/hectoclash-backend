package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
	"github.com/ishu17077/hectoclash-backend/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.POST("/signup", controllers.SignUp())
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.POST("/login", controllers.Login())
	incomingRoutes.GET("/user", controllers.GetUser())

}
