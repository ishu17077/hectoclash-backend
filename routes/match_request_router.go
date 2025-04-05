package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/controllers"
)

func MatchRequest(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/requests", controllers.GetMatchRequests())
	incomingRoutes.PATCH("/request/:request_id", controllers.RespondRequest())
	incomingRoutes.DELETE("/request/:request_id", controllers.RemoveSentRequest())
	incomingRoutes.GET("/requests/sent", controllers.GetSentRequests())
	incomingRoutes.POST("/request/match/:match_id/send/:to_id", controllers.SendMatchRequest())
}
