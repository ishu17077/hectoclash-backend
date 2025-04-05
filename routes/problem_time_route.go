package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func ProblemTimeRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/matches/:match_id/players/:player_id/problems/:problem_number/time-taken", controllers.GetProblemTime())
	incomingRoutes.POST("/matches/:match_id/players/:player_id/problems/:problem_number/time-taken", controllers.CreateProblemTime())
}
