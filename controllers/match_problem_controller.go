package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

var isMatchProblemCollectionChanged bool = false

func GetMatchProblems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		matchId := c.Param("match_id")
		var matchProblems []models.MatchProblem
		defer cancel()
		result, err := matchProblemsCollection.Find(ctx, bson.M{"match_id": matchId})
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding results."})
			cancel()
			return
		}

		if err := result.All(ctx, &matchProblems); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting matches"})
			return
		}

		filter := bson.M{
			"problem_id": bson.M{
				"$in": func() []string {
					var ids []string
					for _, mp := range matchProblems {
						ids = append(ids, *&mp.Problem_id)
					}
					return ids
				}()},
		}
		var problems []models.Problem
		defer cancel()
		result, err = problemCollection.Find(ctx, filter)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding problems"})
			return
		}
		if err := result.All(ctx, &problems); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse the match problems"})
			return
		}
		c.JSON(http.StatusOK, problems)
	}
}

func checkResponse(c *gin.Context, cancel context.CancelFunc) bool {
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
		cancel()
		return true
	}
	return false
}

func GetMatchProblem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		matchId := c.Param("match_id")
		problemNumber, err := strconv.Atoi(c.Param("problem_number")) // Example of using another parameter from query string
		if matchId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please pass Match ID"})
			cancel()
			return
		}
		if err != nil || problemNumber < 1 {
			problemNumber = 1
		}
		filter := bson.M{
			"match_id":       matchId,
			"problem_number": problemNumber,
		}
		var matchProblem models.MatchProblem

		err = matchProblemsCollection.FindOne(ctx, filter).Decode(&matchProblem)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Match Problems not found"})
			return
		}
		var problem models.Problem
		err = problemCollection.FindOne(ctx, bson.M{"problem_id": matchProblem.Problem_id}).Decode(&problem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem not found."})
			return
		}
		c.JSON(http.StatusOK, problem)
	}
}

//?TODO: Future Implementation
// func returnMatchId(matchProblemREST models.MatchProblemREST, c *gin.Context) string {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
// 	if matchProblemREST.Match_problem_id != nil {
// 		err := matchProblemsCollection.FindOne(ctx, bson.M{"match_problem_id": *(matchProblemREST.Match_problem_id)}).Decode(&matchProblemREST)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Match Problem ID passed."})
// 			cancel()
// 			return ""
// 		}
// 		cancel()
// 		return *(matchProblemREST.Match_id)
// 	} else if matchProblemREST.Match_id != nil && (matchProblemREST.Problem_id != nil || matchProblemREST.Problem_number < 1) {
// 		filter := primitive.D{bson.E{Key: "match_id", Value: matchProblemREST.Match_id}}
// 		if matchProblemREST.Problem_id != nil {
// 			filter = append(filter, bson.E{Key: "problem_id", Value: matchProblemREST.Problem_id})
// 		} else if matchProblemREST.Problem_number < 1 {
// 			filter = append(filter, bson.E{Key: "problem_number", Value: matchProblemREST.Problem_number})
// 		} else {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "No Problem ID or Number found."})
// 			cancel()
// 			return ""
// 		}
// 		var matchProblem models.MatchProblem
// 		defer cancel()
// 		err := matchProblemsCollection.FindOne(ctx, filter).Decode(&matchProblem)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Match parameters passed."})
// 			return ""
// 		}
// 		return matchProblem.Match_id
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "No Problem ID or Number found."})
// 		cancel()
// 		return ""
// 	}
// }
