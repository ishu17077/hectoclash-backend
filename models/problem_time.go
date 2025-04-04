package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProblemTime struct {
	ID               primitive.ObjectID `bson:"_id"`
	Match_id         string             `json:"match_id" validate:"required"`
	Problem_id       *string            `json:"problem_id" validate:"required"`
	Problem_number   uint8              `json:"problem_number" validate:"required"`
	Player_id        *string            `json:"player_id" validate:"required"`
	Start_time       time.Time          `json:"start_time"` //! in seconds
	End_time         time.Time          `json:"end_time"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
	Match_problem_id string             `json:"match_problem_id" validate:"required"`
	Problem_time_id  string             `json:"problem_time_id" validate:"required"`
}

type ProblemTimeREST struct {
	Match_id         string    `json:"match_id"`
	Problem_id       *string   `json:"problem_id"`
	Problem_number   uint8     `json:"problem_number"`
	Player_id        *string   `json:"player_id"`
	Start_time       time.Time `json:"start_time"` //! in seconds
	End_time         time.Time `json:"end_time"`
	Match_problem_id string    `json:"match_problem_id"`
	Problem_time_id  string    `json:"problem_time_id"`
}
