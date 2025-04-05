package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func MatchProblemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/match/:match_id/problems", controllers.GetMatchProblems())
	incomingRoutes.GET("/match/:match_id/problem", controllers.GetMatchProblem())
}
