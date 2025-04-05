package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/ishu17077/hectoclash-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

type SignedDetails struct {
	Email    string
	Username string
	Uid      string
	jwt.RegisteredClaims
}

var secretKey string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, username string, uid string) (signedToken string, refreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		Username:   username,
		Uid:        uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))),
			//? Recommended time is 30 minutes but since we are building for test 24 hours is okay
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(168))),
			//? refresh token shoul be around 24 hours in prod
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updated_at})

	upsert := true

	filter := bson.M{"user_id": userId}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
	return
}

func ValiateToken(signedToken string) (claims *SignedDetails, msg string) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		msg = fmt.Sprintf("Error parsing token: %v", err)
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok || !token.Valid {
		msg = "The token is invalid"
		return
	}

	//Token is valid
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		msg = fmt.Sprintf("Token expired")
		return
	}
	return claims, msg
}
