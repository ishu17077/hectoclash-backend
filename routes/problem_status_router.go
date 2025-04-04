package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func ProblemStatusRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/match/problem-status", controllers.GetProblemStatus())
	incomingRoutes.GET("/match/problem-statuses", controllers.GetProblemStatuses())
	incomingRoutes.POST("/matches/:match_id/players/:player_id/problems/:problem_number/problem_status", controllers.CreateProblemStatus())
}
