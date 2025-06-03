package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Match struct {
	ID                   primitive.ObjectID `bson:"_id"`
	Match_code           uint32             `json:"match_code" validate:"required,min=100000,max=999999"`
	Player_ids           []string           `json:"player_ids" validate:"required"`
	Problems             uint8              `json:"problems" validate:"required,min=1"`
	Created_at           time.Time          `json:"created_at"`
	Updated_at           time.Time          `json:"updated_at"`
	Started_at           time.Time          `json:"started_at"`
	Match_type           string             `json:"match_type" validate:"required,oneof=SOLO DUAL"`
	Was_interrupted      bool               `json:"was_interrupted" bson:"was_interrupted" default:"false"`
	Time_elapsed         time.Duration      `json:"time_elapsed"`
	Is_started           bool               `json:"is_started"`
	Is_private           bool               `json:"is_private"`
	Created_by_player_id string             `json:"created_by_player_id" validate:"required"`
	Match_id             string             `json:"match_id" validate:"required"`
	Match_duration       time.Duration      `json:"match_duration" validate:"required"`
	Viewers              uint16             `json:"viewers" validate:"min=0"`
	Is_ended             bool               `json:"is_ended"`
}

type MatchREST struct {
	Match_code           uint32        `json:"match_code"`
	Problems             uint8         `json:"problems"`
	Created_by_player_id string        `json:"created_by_player_id"`
	Player_ids           []string      `json:"player_ids" validate:"required"`
	Match_id             string        `json:"match_id"`
	Match_type           string        `json:"match_type" validate:"required,oneof=SOLO DUAL"`
	Match_duration       time.Duration `json:"match_duration"`
	Is_private           bool          `json:"is_private"`
	Is_ended             bool          `json:"is_ended"`
}
