package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/database"
	"github.com/ishu17077/hectoclash-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var problemCollection *mongo.Collection = database.OpenCollection(database.Client, "problems")

func GetProblems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var problems []models.Problem
		defer cancel()
		result, err := problemCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching problems from database"})
			return
		}
		if err := result.All(ctx, &problems); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listing problems from database"})
			return
		}
		c.JSON(http.StatusOK, problems)
	}
}

func GetProblem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var problemId string = c.Params.ByName("problem_id")
		var problem models.Problem
		defer cancel()
		err := problemCollection.FindOne(ctx, bson.M{"problem_id": problemId}).Decode(&problem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching problems from database"})
			return
		}
		c.JSON(http.StatusOK, problem)
	}
}
