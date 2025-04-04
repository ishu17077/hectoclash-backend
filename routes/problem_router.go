package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func ProblemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/problems", controllers.GetProblems())
	incomingRoutes.GET("/problems/:problem_id", controllers.GetProblem())
	// incomingRoutes.POST("/problems",)
	// incomingRoutes.PATCH("/problems/:problem_id",)
}
