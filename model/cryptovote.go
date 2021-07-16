package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CryptoVote struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Symbol       string             `json:"symbol" bson:"symbol"`
	Qtd_Upvote   int                `json:"qtd_upvote" bson:"qtd_upvote"`
	Qtd_Downvote int                `json:"qtd_downvote" bson:"qtd_downvote"`
}
