package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func MatchProblemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/match/problems", controllers.GetMatchProblems())
	incomingRoutes.GET("/matches/problem", controllers.GetMatchProblem())
}
