package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func MatchRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/matches", controllers.GetMatches())
	incomingRoutes.GET("/match/:match_id", controllers.GetMatch())
	//! Args should be passed as problem-time like ../problem-time?problem=1
	incomingRoutes.POST("/match/:match_id/start-match", controllers.StartMatch())
	incomingRoutes.POST("/match", controllers.CreateMatch())

}
