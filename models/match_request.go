package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MatchRequest struct {
	ID               primitive.ObjectID `bson:"_id"`
	From_id          *string            `json:"from_id" validate:"required"`
	To_id            *string            `json:"to_id" validate:"required"`
	Status           bool               `json:"status" validate:"required,eq=ACCEPTED||eq=IGNORED||eq=NOTSEEN"`
	Match_request_id string             `json:"match_request_id" validate:"required"`
}

type MatchRequestREST struct {
	From_id          *string `json:"from_id"`
	To_id            *string `json:"to_id"`
	Match_request_id *string `json:"match_request_id"`
}
