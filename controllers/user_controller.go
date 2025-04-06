package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ishu17077/hectoclash-backend/database"
	"github.com/ishu17077/hectoclash-backend/helpers"
	"github.com/ishu17077/hectoclash-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

// ! Internal Function to query
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
				{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}, //? data we used and $data is we accessing at project stage
			}}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: bson.D{{Key: "$size", Value: "$data"}}},
			{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}},
		}}}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			groupStage,
			projectStage,
		})
		defer cancel()
		//either pass an error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var allUsers []bson.M
		if err := result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)

		}
		//ideally want to return all the users based on verious query parameter
		c.JSON(http.StatusOK, allUsers)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := c.GetString("uid")
		fmt.Print(userId)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding the user"})
			return
		}
		var userREST models.UserREST = models.UserREST{
			Username: user.Username,
			Avatar:   user.Avatar,
			Password: user.Password,
			Email:    user.Email,
			Phone:    user.Phone,
			User_id: user.User_id,
		}
		c.JSON(http.StatusOK, gin.H{
			"username": userREST.Username,
			"email":    userREST.Email,
			"phone":    userREST.Phone,
			"avatar":   userREST.Avatar,
			"user_id":	user.User_id,

		})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var userREST models.UserREST
		var user models.User

		//* Convert the JSON data coming from postman to something go lang would understand
		if err := c.BindJSON(&userREST); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing json"})
			cancel()
			return
		}
		//* get some extra details for user object like created_at, updated_at,ID

		user.Email = userREST.Email
		user.Phone = userREST.Phone
		user.Password = userREST.Password
		user.Username = userREST.Username
		user.Avatar = userREST.Avatar
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		user.Player_id = user.User_id

		//* validate the above data based on user struct
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			cancel()
			return
		}
		//* check if email is already used by another user
		//* check if phone no or mail is already used by another user
		defer cancel()
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": *(&user.Email)})

		//! let's see if this works instead of querying two times

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking email or phone"})
			log.Panic(err)

		}
		//* hash password
		password := HashPassword(*user.Password)
		user.Password = &password

		if count > 0 {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Email or Phone already in use"})
			return
		}

		//* generate token and refresh tokens (generate all tokens from helper)

		token, refeshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.Username, user.User_id)
		user.Token = &token
		user.Refresh_token = &refeshToken
		//* if all okay, insert this new user to user collection
		_, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		//* return statusOk and send the result back
		c.JSON(http.StatusOK, user)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var userREST models.UserREST
		var foundUser models.User
		//* Convert the Login JSON data coming from postman to something go lang would understand
		if err := c.BindJSON(&userREST); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			defer cancel()
			return
		}
		//* find a user with that email and check if exist

		var filter bson.M

		if userREST.Email != nil {
			filter = bson.M{"email": *userREST.Email}
		} else if userREST.Phone != nil {
			filter = bson.M{"phone": *userREST.Phone}
		}

		err := userCollection.FindOne(ctx, filter).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found!"})
			return
		}
		//* if user exists, verify the password
		passwordIsValid, msg := verifyPassword(*foundUser.Password, *userREST.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}
		//* if all goes well, then you'll generate tokens
		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.Username, *&foundUser.User_id)

		//* update tokens - token and refresh tokens
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		//* return StatusOK
		c.JSON(http.StatusOK, gin.H{
			"username":      foundUser.Username,
			"email":         foundUser.Email,
			"avatar":        foundUser.Avatar,
			"phone":         foundUser.Phone,
			"token":         foundUser.Token,
			"refresh_token": foundUser.Refresh_token,
			"user_id":		foundUser.User_id,
		})
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func verifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	check := false
	msg := "Password Incorrect"
	if err == nil {
		check = true
		msg = fmt.Sprintf("Login Success!")
	}
	return check, msg
}
