package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPlayerScorecard() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		matchId := c.Param("match_id")
		playerId := c.GetString("uid")
		filter := bson.M{
			"match_id":  matchId,
			"player_id": playerId,
		}
		var playerScorecard models.PlayerScorecard

		defer cancel()
		err := playerScorecardCollection.FindOne(ctx, filter).Decode(&playerScorecard)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Score card for the following is not available."})
		}

		// if err != nil {
		// 	var playerScorecard models.PlayerScorecard
		// 	filterMatchCol := bson.M{
		// 		"match_id":             matchId,
		// 		"created_by_player_id": playerId,
		// 	}
		// 	defer cancel()
		// 	var match models.Match
		// 	err = matchesCollection.FindOne(ctx, filterMatchCol).Decode(&match)
		// 	if err != nil {
		// 		c.JSON(http.StatusBadRequest, gin.H{"error": "No match with the match id found."})
		// 		return
		// 	}
		// 	var problem_statuses []models.ProblemStatus
		// 	result, err := problemStatusCollection.Find(ctx, filter)

		// 	if err != nil {
		// 		if result.RemainingBatchLength() == 0 {

		// 		}
		// 	}
		// }
		c.JSON(http.StatusOK, playerScorecard)
	}
}
