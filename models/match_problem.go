package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MatchProblem struct {
	ID               primitive.ObjectID `bson:"_id"`
	Match_id         string             `json:"match_id" validate:"required"`
	Problem_id       string             `json:"problem_id" validate:"required"`
	Problem_number   uint8              `json:"problem_number" validate:"required"`
	Match_problem_id string             `json:"match_problem_id" validate:"required"` //? Internally handled id of MatchProblem class
	Created_at       time.Time          `json:"Created_at"`
	Updated_at       time.Time          `json:"Updated_at"`
}

type MatchProblemREST struct {
	Match_id         *string `json:"match_id"`
	Problem_id       *string `json:"problem_id"`
	Problem_number   uint8   `json:"problem_number"`
	Match_problem_id *string `json:"match_problem_id"`
}

type MatchProblemNoProb struct {
	ID               primitive.ObjectID `bson:"_id"`
	Match_id         string             `json:"match_id" validate:"required"`
	Problem_id       string             `json:"problem_id" validate:"required"`
	Problem          Problem            `json:"problem" validate:"required"`
	Problem_number   uint8              `json:"problem_number"`
	Match_problem_id string             `json:"match_problem_id" validate:"required"` //? Internally handled id of MatchProblem class
	Created_at       time.Time          `json:"Created_at"`
	Updated_at       time.Time          `json:"Updated_at"`
}
