package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerScorecard struct {
	ID                  primitive.ObjectID `bson:"_id"`
	Player_scorecard_id string             `json:"player_scorecard_id" validate:"required"`
	Player_id           string             `json:"player_id" validate:"required"`
	Match_id            string             `json:"match_id" validate:"required"`
	Start_time          time.Time          `json:"start_time"`
	End_time            time.Time          `json:"end_time"`
	Total_points        uint16             `json:"total_points" validate:"required"`
	Remarks             string             `json:"remarks" validate:"eq=EXCELLENT||eq=GOOD||eq=FAIR||eq=IMPROVNEEDED||eq="`
	Insights            string             `json:"insights"`
	Attempt_ended       bool               `json:"attempt_ended"`
}

type PlayerScorecardREST struct {
	Player_scorecard_id *string   `json:"player_scorecard_id"`
	Player_id           *string   `json:"player_id"`
	Match_id            *string   `json:"match_id"`
	Start_time          time.Time `json:"start_time"`
	End_time            time.Time `json:"end_time"`
	Attempt_ended       bool      `json:"attempt_ended"`
}
