package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CryptoVote struct {
	Id           primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name"`
	Symbol       string             `json:"symbol,omitempty" bson:"symbol"`
	Qtd_Upvote   int                `json:"qtd_upvote,omitempty" bson:"qtd_upvote"`
	Qtd_Downvote int                `json:"qtd_downvote,omitempty" bson:"qtd_downvote"`
	Sum          int                `json:"sum,omitempty" bson:"sum"`
	SumAbsolute  int                `json:"sum_absolute,omitempty" bson:"sum_absolute"`
}

type FilterCryptoVote struct {
	Id     primitive.ObjectID `json:"-" bson:"-"`
	Name   string             `json:"name,omitempty" bson:"name"`
	Symbol string             `json:"symbol,omitempty" bson:"symbol"`
}

type SumaryVote struct {
	Crypto      CryptoVote `json:"-" bson:"-"`
	SumToken    string     `json:"sum_token" bson:"sum_token"`
	SumType     string     `json:"sum_type" bson:"sum_type"`
	SumAbsolute int        `json:"sum_absolute" bson:"sum_absolute"`
}
