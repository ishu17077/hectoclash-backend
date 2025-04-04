package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProblemStatus struct {
	ID                primitive.ObjectID `bson:"_id"`
	Player_id         string             `json:"player_id" validate:"required"`
	Match_problem_id  string             `json:"match_problem_id" validate:"required"` //? Internally handled id of MatchProblem class
	Match_id          string             `json:"match_id" validate:"required"`
	Points            uint16             `json:"points" validate:"required"`
	Problem_number    uint8              `json:"problem_number"`
	Is_solved         bool               `json:"is_solved"`
	Created_at        time.Time          `json:"created_at"`
	Updated_at        time.Time          `json:"updated_at"`
	Problem_status_id string             `json:"problem_status_id" validate:"required"`
}

type ProblemStatusREST struct {
	Player_id         *string `json:"player_id"`
	Match_problem_id  *string `json:"match_problem_id"` //? Internally handled id of MatchProblem class
	Match_id          *string `json:"match_id"`
	Points            uint16 `json:"points"`
	Problem_number    uint8  `json:"problem_number"`
	Is_solved         bool   `json:"is_solved"`
	Problem_status_id *string `json:"problem_status_id" validate:"required"`
}
