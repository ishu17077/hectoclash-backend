package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Username      *string            `json:"username" validate:"required,min=2,max=100"`
	Password      *string            `json:"password" validate:"required,min=6"`
	Email         *string            `json:"email" validate:"email,required"`
	Avatar        *string            `json:"avatar"`
	Phone         *string            `json:"phone"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id" validate:"required"`
	Player_id     string             `json:"player_id" validate:"required"`
}

type UserREST struct {
	Username *string `json:"username" validate:"required,min=2,max=100"`
	Password *string `json:"password" validate:"required,min=6"`
	Email    *string `json:"email" validate:"email,required"`
	Avatar   *string `json:"avatar"`
	Phone    *string `json:"phone"`
	User_id  string  `json:"user_id"`
}
