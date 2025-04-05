package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func ProblemStatusRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/match/:match_id/problem/:problem_number/problem-status", controllers.GetProblemStatus())
	incomingRoutes.GET("/match/:match_id/problem-statuses", controllers.GetProblemStatuses())
	incomingRoutes.POST("/match/:match_id/:problem_number/problem-status", controllers.CreateProblemStatus())
}
