package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/user/:user_id", controllers.GetUser())
	incomingRoutes.POST("/user/signup", controllers.SignUp())
	incomingRoutes.POST("/user/login", controllers.Login())
}
