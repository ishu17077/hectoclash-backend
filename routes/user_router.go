package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/user", controllers.GetUser())
	incomingRoutes.POST("/signup", controllers.SignUp())
	incomingRoutes.POST("/login", controllers.Login())
}
