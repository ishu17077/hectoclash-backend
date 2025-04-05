package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProblemStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		playerId := c.Keys["uid"]
		fmt.Print(playerId)
		matchId := c.Param("match_id")
		problemNumber := c.Param("problem_number")

		filter := bson.M{
			"match_id":       matchId,
			"problem_number": problemNumber,
			"player_id":      playerId,
		}

		var problemStatus models.ProblemStatus

		defer cancel()
		err := problemStatusCollection.FindOne(ctx, filter).Decode(&problemStatus)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem Status not found"})
			return
		}
		c.JSON(http.StatusOK, problemStatus)
	}
}
func GetProblemStatuses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		playerId := c.Keys["uid"]
		matchId := c.Param("match_id")

		filter := bson.M{
			"match_id":  matchId,
			"player_id": playerId,
		}

		var problemStatuses []models.ProblemStatus
		defer cancel()

		result, err := problemStatusCollection.Find(ctx, filter)
		if err = result.All(ctx, &problemStatuses); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching problem statuses"})
			return
		}
		c.JSON(http.StatusOK, problemStatuses)
	}
}

func CreateProblemStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		playerId := c.GetString("uid")
		var problemStatusREST models.ProblemStatusREST

		if err := c.BindJSON(&problemStatusREST); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			cancel()
			return
		}

		var filter bson.M
		if problemStatusREST.Match_id != nil && problemStatusREST.Problem_number > 0 {
			filter = bson.M{
				"match_id":       *(problemStatusREST.Match_id),
				"player_id":      playerId,
				"problem_number": problemStatusREST.Problem_number,
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please pass appropriate values."})
			cancel()
			return
		}

		var matchProblem models.MatchProblem

		var problemStatus models.ProblemStatus
		defer cancel()
		err := problemStatusCollection.FindOne(ctx, filter).Decode(&problemStatus)

		if err == nil {
			c.JSON(http.StatusOK, problemStatus)
			return
		} else {
			filter = bson.M{
				"match_id":       *(problemStatusREST.Match_id),
				"problem_number": problemStatusREST.Problem_number,
			}

			defer cancel()
			err := matchProblemsCollection.FindOne(ctx, filter).Decode(&matchProblem)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching Match Problem"})
				return
			}

			currentTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			problemStatus = models.ProblemStatus{
				ID:               primitive.NewObjectID(),
				Player_id:        playerId,
				Match_id:         matchProblem.Match_id,
				Points:           problemStatus.Points,
				Problem_number:   uint8(matchProblem.Problem_number),
				Match_problem_id: matchProblem.Problem_id,
				Attempted:        true,
				Is_solved:        false,
				Created_at:       currentTime,
				Updated_at:       currentTime,
			}
			//? Implement a time up

			problemStatus.Problem_status_id = problemStatus.ID.Hex()

			if validationErr := validate.Struct(&problemStatus); validationErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error Processing your data"})
				return
			}
			defer cancel()
			_, err = problemStatusCollection.InsertOne(ctx, problemStatus)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting new problem status"})
				return
			}

			filter := bson.M{
				"match_id":       matchProblem.Match_id,
				"player_id":      playerId,
				"problem_number": matchProblem.Problem_number,
			}
			update := bson.M{
				"$inc": bson.M{
					"points": problemStatus.Points, // Increment the points by the value in problemStatus.Points
				},
			}

			_, err = playerScorecardCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating score card"})
				return
			}

			c.JSON(http.StatusOK, problemStatus)
			return
		}
	}

}
