package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/database"
	"github.com/ishu17077/hectoclash-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

var isMatchRequestCollectionChanged bool = true

var matchRequestCollection = database.OpenCollection(database.Client, "match_requests")

func GetMatchRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerId := c.GetString("uid")
		var matchRequests []models.MatchRequest
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		result, err := matchRequestCollection.Find(ctx, bson.M{"to_id": playerId, "status": "NOTSEEN"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding"})
		}
		if result.RemainingBatchLength() <= 0 {
			c.JSON(http.StatusOK, "[]")
			return
		}
		if err := result.All(ctx, &matchRequests); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error "})
			return
		}
		c.JSON(http.StatusOK, matchRequests)
	}
}

func GetSentRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerId := c.GetString("uid")
		var matchRequests []models.MatchRequest
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		result, err := matchRequestCollection.Find(ctx, bson.M{"from_id": playerId, "status": "NOTSEEN"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding"})
		}
		if result.RemainingBatchLength() <= 0 {
			c.JSON(http.StatusOK, "[]")
			return
		}
		if err := result.All(ctx, &matchRequests); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error "})
			return
		}
		c.JSON(http.StatusOK, matchRequests)
	}
}

func RemoveSentRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerId := c.GetString("uid")
		match_request_id := c.Param("match_request_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		result, err := matchRequestCollection.DeleteOne(ctx, bson.M{"match_request_id": match_request_id, "from_id": playerId})
		isMatchRequestCollectionChanged = true
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding any request"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func RespondRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerId := c.GetString("uid")
		match_request_id := c.Param("match_request_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		result, err := matchRequestCollection.DeleteOne(ctx, bson.M{"match_request_id": match_request_id, "to_id": playerId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding any request"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func ptr(s string) *string {
	return &s
}

func SendMatchRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerId := c.GetString("uid")
		toId := c.Param("to_id")
		match_id := c.Param("match_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var matchRequest models.MatchRequest = models.MatchRequest{
			Status:   ptr("NOTSEEN"),
			From_id:  &playerId,
			To_id:    &toId,
			Match_id: &match_id,
		}
		matchRequest.Match_request_id = matchRequest.ID.Hex()
		defer cancel()
		_, err := matchRequestCollection.InsertOne(ctx, matchRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding any request"})
			return
		}

		c.JSON(http.StatusOK, matchRequest)
	}
}

//TODO: Match_req_acceptedl

func UpdateMessageRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerId := c.GetString("uid")
		response := c.Query("response")
		matchRequestId := c.Param("match_request_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var matchRequest models.MatchRequest
		defer cancel()
		if response == "IGNORED" || response == "ACCEPTED" || response == "NOTSEEN" {
			err := matchRequestCollection.FindOneAndUpdate(ctx, bson.M{"match_request_id": matchRequestId, "to_id": playerId}, bson.M{"status": response}).Decode(&matchRequest)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing to match request."})
			}
		}
		if response == "ACCEPTED" {
			defer cancel()
			var match models.Match
			err := matchesCollection.FindOneAndUpdate(
				ctx,
				bson.M{"match_id": matchRequest.Match_id},
				bson.M{"$addToSet": bson.M{"player_ids": playerId}},
			).Decode(&match)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing to match."})
			}
		}
		c.JSON(http.StatusOK, matchRequest)
	}
}
