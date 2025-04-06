package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Problem struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Problem_components []string           `json:"problem_components" validate:"required"`
	Answer             int32              `json:"answer" validate:"required"`
	Right_sequence     []string           `json:"right_sequence"`
	Operators          []string           `json:"operators"`
	Points             uint16             `json:"points" validate:"required"`
	Solve_time         time.Duration      `json:"solve_time" validate:"required"`
	Min_points         uint16             `json:"min_points"`
	Hint               string             `json:"hint"`
	Created_at         time.Time          `json:"created_at"`
	Updated_at         time.Time          `json:"updated_at"`
	Problem_id         string             `json:"problem_id" validate:"required"`
}

type problemCollectioREST struct {
	Problem_id string `json:"problem_id" validate:"required"`
}
