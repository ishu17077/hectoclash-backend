package controllers

import (
	"context"

	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ishu17077/hectoclash-backend/database"
	"github.com/ishu17077/hectoclash-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var matchesCollection *mongo.Collection = database.OpenCollection(database.Client, "matches")

var matchProblemsCollection *mongo.Collection = database.OpenCollection(database.Client, "match_problems")
var problemTimeCollection *mongo.Collection = database.OpenCollection(database.Client, "problem_time")
var problemStatusCollection *mongo.Collection = database.OpenCollection(database.Client, "problem_status")
var playerScorecardCollection *mongo.Collection = database.OpenCollection(database.Client, "player_scorecard")
var validate *validator.Validate = validator.New()

var isMatchCollectionChanged bool = true

func GetMatches() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordsPerPage, err := strconv.Atoi(c.Query("recordsPerPage"))
		if err != nil || recordsPerPage < 1 || recordsPerPage > 150 {
			recordsPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordsPerPage
		startIndex, _ = strconv.Atoi(c.Query("startIndex"))
		matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "is_ended", Value: false}}}}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
				{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}, //? data we used and $data is we accessing at project stage
			}}}
		projectStage := bson.D{{
			Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "total_count", Value: 1},
				{Key: "matches", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordsPerPage}}}},
			}}}
		defer cancel()
		result, err := matchesCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while listing matches."})
			return
		}
		var allMatches []bson.M

		if err := result.All(ctx, &allMatches); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while parsing the matches."})
			return
		}
		c.JSON(http.StatusOK, allMatches[0]) //? allMatches[0] could be returned

	}

}

func GetMatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var matchId string = c.Params.ByName("match_id")
		var match models.Match
		defer cancel()
		err := matchesCollection.FindOne(ctx, bson.M{"match_id": matchId}).Decode(&match)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, match)
	}
}

func CreateMatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var matchREST models.MatchREST
		playerId := c.GetString("uid")
		var match models.Match
		if err := c.BindJSON(&matchREST); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			cancel()
			return
		}
		if matchREST.Problems < 1 || matchREST.Problems > 20 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error Bad Request"})
			cancel()
			return
		}
		defer cancel()
		// result, err := matchesCollection.Find(ctx, bson.M{"created_by_player_id": playerId, "is_ended": false})

		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }
		// var allMatches []models.Match
		// if result.RemainingBatchLength() > 0 {
		// 	if err := result.All(ctx, &allMatches); err == nil {

		// 		matchStage := bson.D{
		// 			{Key: "$match", Value: bson.D{{Key: "match_id", Value: allMatches[0].Match_id}}},
		// 		}
		// 		aggregationStage := bson.D{
		// 			{Key: "$lookup", Value: bson.D{
		// 				{Key: "from", Value: "problems"},
		// 				{Key: "localField", Value: "problem_id"},
		// 				{Key: "foreignField", Value: "problem_id"},
		// 				{Key: "as", Value: "problems"},
		// 			}}}
		// 		defer cancel()
		// 		res, err := matchesCollection.Aggregate(ctx, mongo.Pipeline{matchStage, aggregationStage})
		// 		if err == nil {
		// 			var existingMatchProblems []models.MatchProblem
		// 			result, err = matchProblemsCollection.Find(ctx, bson.M{"match_id": allMatches[0].Match_id})
		// 			if err := result.All(ctx, &existingMatchProblems); err == nil {
		// 				for matchProblems in

		// 				var jsony []bson.M
		// 				if err := res.All(ctx, &jsony); err == nil {
		// 					c.JSON(http.StatusOK, jsony)
		// 					return
		// 				}
		// 			}

		// 		}

		match.Problems = matchREST.Problems
		match.ID = primitive.NewObjectID()
		match.Match_code = uint32(rand.Int31n(899999) + 100000)
		//TODO: To be implemented: random number check in database
		// Replace 10 and 20 with actual values for 'a' and 'b'
		match.Match_id = match.ID.Hex()
		match.Match_type = matchREST.Match_type
		match.Is_started = false
		match.Created_by_player_id = playerId
		match.Is_ended = false
		match.Viewers = 0
		match.Is_private = false
		match.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		match.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		match.Player_ids = []string{playerId}
		// Validate the match struct
		// if validationErr := validate.Struct(&match); validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed: " + validationErr.Error()})
		// 	cancel()
		// 	return
		// }
		defer cancel()
		_, insertErr := matchesCollection.InsertOne(ctx, match)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}
		// _, err := InitializeMatchComponents(match)
		// if err != nil {
		// 	log.Fatal(err)
		//
		if match.Match_type == "SOLO" {
			startMatchHandler := StartMatch()
			c.Params = append(c.Params, gin.Param{Key: "match_id", Value: match.Match_id})
			startMatchHandler(c)

		} else {
			// c.JSON(http.StatusOK, result)
		}

		return
	}
}

func StartMatch() gin.HandlerFunc {
	return func(c *gin.Context) {

		matchId := c.Param("match_id")

		playerId := c.GetString("uid")
		var updateObj bson.D

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var match models.Match

		currentTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = bson.D{
			{Key: "is_started", Value: true},
			{Key: "is_ended", Value: false},
			{Key: "started_at", Value: currentTime},
			{Key: "updated_at", Value: currentTime},
		}

		// Upsert option
		upsert := true
		opt := options.FindOneAndUpdateOptions{
			Upsert: &upsert,
		}

		filter := bson.M{"created_by_player_id": playerId, "is_ended": false}

		defer cancel()
		err := matchesCollection.FindOneAndUpdate(ctx, filter, bson.D{
			{Key: "$set", Value: updateObj},
		}, &opt).Decode(&match)
		if match.Is_ended == false {
			match.Is_started = true
			match.Started_at = currentTime
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The Match has already ended."})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer cancel()
		var existingMatchProblems []models.MatchProblem
		result, err := matchProblemsCollection.Find(ctx, bson.M{"match_id": matchId})
		defer cancel()
		err = result.All(ctx, &existingMatchProblems)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching match problems."})
			return
		}

		if len(existingMatchProblems) <= 0 && match.Is_ended == false {
			var matchProblems []models.MatchProblem
			var problems []models.Problem

			aggregation := []bson.M{
				{"$sample": bson.M{"size": int(match.Problems)}},
			}
			defer cancel()
			result, err := problemCollection.Aggregate(ctx, aggregation)
			if err != nil {
				log.Fatal(err)
			}
			if err := result.All(ctx, &problems); err != nil {
				log.Fatal(err)
			}
			matchProblemsInterface := []interface{}{}
			for i := 1; i <= int(match.Problems); i++ {
				matchProblem := models.MatchProblem{
					ID:             primitive.NewObjectID(),
					Match_id:       match.Match_id,
					Problem_id:     problems[i-1].Problem_id,
					Problem_number: uint8(i),
				}
				matchProblem.Match_problem_id = matchProblem.ID.Hex()
				matchProblem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
				matchProblem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

				if validationErr := validate.Struct(&matchProblem); validationErr != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Please make sure you have this match registered."})
					return
				}
				matchProblems = append(matchProblems, matchProblem)
				matchProblemsInterface = append(matchProblemsInterface, matchProblem)
			}

			_, err = matchProblemsCollection.InsertMany(ctx, matchProblemsInterface)
			if err != nil {
				log.Print(err)

			}
			isMatchProblemCollectionChanged = true
			defer cancel()

			var playerScorecard models.PlayerScorecard = models.PlayerScorecard{
				ID:            primitive.NewObjectID(),
				Player_id:     playerId,
				Match_id:      matchId,
				Start_time:    currentTime,
				Total_points:  0,
				Attempt_ended: false,
			}
			playerScorecard.Player_scorecard_id = playerScorecard.ID.Hex()

			_, err = playerScorecardCollection.InsertOne(ctx, playerScorecard)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding player scorecard"})
				return
			}

			routineUpdateUnusedMatches(matchId, playerId)

			var ids []string
			for _, matchProblem := range matchProblems {
				ids = append(ids, *&matchProblem.Match_problem_id)
			}
			if len(ids) == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": "No match problems found."})
				return
			}
			filter := bson.D{{Key: "match_problem_id", Value: bson.D{{Key: "$in", Value: ids}}}}

			defer cancel()
			matchStage := bson.D{{Key: "$match", Value: filter}}
			leftJoinAggregation := bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "problems"},
				{Key: "localField", Value: "problem_id"},
				{Key: "foreignField", Value: "problem_id"},
				{Key: "as", Value: "problems"},
				// $match value

			}}}
			res, err := matchProblemsCollection.Aggregate(ctx, mongo.Pipeline{matchStage, leftJoinAggregation})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			var matchProblemsNoProb []bson.M
			_ = res.All(ctx, &matchProblemsNoProb)
			c.JSON(http.StatusOK, matchProblemsNoProb)
		} else {
			var ids []string
			for _, matchProblem := range existingMatchProblems {
				ids = append(ids, *&matchProblem.Match_problem_id)
			}
			if len(ids) == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": "No match problems found."})
				return
			}

			filter := bson.D{{Key: "match_problem_id", Value: bson.D{{Key: "$in", Value: ids}}}}

			defer cancel()
			matchStage := bson.D{{Key: "$match", Value: filter}}
			leftJoinAggregation := bson.D{{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "problems"},
				bson.E{Key: "localField", Value: "problem_id"},
				bson.E{Key: "foreignField", Value: "problem_id"},
				bson.E{Key: "as", Value: "problems"},
			}}}
			res, err := matchProblemsCollection.Aggregate(ctx, mongo.Pipeline{matchStage, leftJoinAggregation})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing items."})
				return
			}

			var matchProblemsNoProb []bson.M
			if err := res.All(ctx, &matchProblemsNoProb); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding match problems."})
				return
			}

			c.JSON(http.StatusOK, matchProblemsNoProb)
		}

	}
}

func routineUpdateUnusedMatches(matchId string, playerId string) {
	go func(matchId string, playerId string) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var updateObj bson.D = bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "is_started", Value: true},
				{Key: "is_ended", Value: true},
			}}}
		time.Sleep(20 * time.Minute)
		_, err := matchesCollection.UpdateOne(ctx, bson.M{"match_id": matchId}, updateObj)
		if err != nil {
			log.Printf("Error updating is_ended for match_id %s: %v", matchId, err)
		}

		currentTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		defer cancel()

		updateObj = bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "attempt_ended", Value: true},
				{Key: "end_time", Value: bson.D{
					{Key: "$cond", Value: bson.D{
						{Key: "if", Value: bson.E{Key: "$eq", Value: []any{"$end_time", nil}}},
						{Key: "then", Value: currentTime},
						{Key: "else", Value: "$end_time"},
					}},
				}},
			}},
		}

		_, err = playerScorecardCollection.UpdateOne(ctx, bson.M{"match_id": matchId, "player_id": playerId}, updateObj)

		if err != nil {
			log.Printf("Error updating is_ended for match_id %s: %v", matchId, err)
		}

	}(matchId, playerId)
}
func AddPauseToCurrentMatches() bool {
	var ctx, cancel = context.WithTimeout(context.TODO(), time.Minute)
	var allMatchesToBePaused []models.Match
	filter := bson.D{
		{Key: "is_started", Value: true},
		{Key: "is_ended", Value: false},
	}
	upsert := true
	// var updateObj bson.D
	opt := options.FindOneAndUpdateOptions{
		Upsert: &upsert,
	}

	// updateObj = append(updateObj, bson.E{"was_interrupted", true})
	// updateObj = append(updateObj, bson.E{"time_elapsed", })

	var currentTime = time.Now()
	defer cancel()
	result, err := matchesCollection.Find(ctx, filter)

	if err != nil {
		log.Printf("Something is wrong finding the matches: %s", err.Error())
		cancel()
		return false
	}

	if err = result.All(ctx, &allMatchesToBePaused); err != nil {
		log.Printf("Something is wrong parsing the matches: %s", err.Error())
		cancel()
		return false
	}
	defer cancel()
	for _, matchToBePaused := range allMatchesToBePaused {
		if currentTime.Unix()-matchToBePaused.Started_at.Unix() >= 1200 { //? 1200 seconds or 20 mins
			matchToBePaused.Is_ended = true
			matchToBePaused.Time_elapsed = 20 * time.Minute
			matchToBePaused.Was_interrupted = false

			matchesCollection.FindOneAndUpdate(ctx, bson.M{"match_id": matchToBePaused.Match_id}, matchToBePaused, &opt)
			continue
		}

		matchToBePaused.Was_interrupted = true
		matchToBePaused.Time_elapsed = time.Duration(currentTime.Unix()-matchToBePaused.Started_at.Unix()) * time.Second
		matchesCollection.FindOneAndUpdate(ctx, bson.M{"match_id": matchToBePaused.Match_id}, matchToBePaused, &opt)

	}

	return true
}
