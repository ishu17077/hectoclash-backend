package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func PlayerScorecardRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/matches/:match_id/players/:player_id/scorecard", controllers.GetPlayerScorecard())
}
