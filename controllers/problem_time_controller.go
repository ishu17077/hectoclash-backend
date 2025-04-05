package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProblemTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		matchId := c.Param("match_id")
		playerId := c.Param("player_id")
		problemNumber, err := strconv.Atoi(c.Param("problem_number"))
		if err != nil || problemNumber < 1 || matchId == "" || playerId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			cancel()
			return
		}
		filter := bson.M{
			"match_id":       matchId,
			"player_id":      playerId,
			"problem_number": problemNumber,
		}
		defer cancel()
		result, err := problemTimeCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while listing problem times"})
			return
		}
		var problemTimes []models.ProblemTime
		if err := result.All(ctx, &problemTimes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while parsing problem times"})
		}
		c.JSON(http.StatusOK, problemTimes)
	}
}

func CreateProblemTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		matchId := c.Param("match_id")
		playerId := c.Param("player_id")
		problemNumber, err := strconv.Atoi(c.Param("problem_number"))
		if err != nil || problemNumber < 1 || matchId == "" || playerId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			cancel()
			return
		}
		var problemTimeREST models.ProblemTimeREST
		if err := c.BindJSON(&problemTimeREST); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			cancel()
			return
		}
		if problemTimeREST.Problem_number < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please pass match_id and problem number"})
			cancel()
			return
		}

		filterMatchProblem := bson.M{
			"match_id":       matchId,
			"problem_number": problemNumber,
		}
		defer cancel()
		var matchProblem models.MatchProblem
		var erre = matchProblemsCollection.FindOne(ctx, filterMatchProblem).Decode(&matchProblem)
		if erre != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while listing problem times"})
			return
		}
		var problemTime models.ProblemTime = models.ProblemTime{
			ID:               primitive.NewObjectID(),
			Match_id:         matchId,
			Problem_id:       matchProblem.Problem_id,
			Problem_number:   matchProblem.Problem_number,
			Player_id:        playerId,
			Start_time:       problemTimeREST.Start_time,
			End_time:         problemTimeREST.End_time,
			Match_problem_id: matchProblem.Match_id,
		}
		problemTime.Problem_time_id = problemTime.ID.Hex()
		problemTime.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		problemTime.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		defer cancel()
		_, err = problemTimeCollection.InsertOne(ctx, problemTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, problemTime)
	}
}
