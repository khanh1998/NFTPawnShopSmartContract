package model

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bid struct {
	UUID               primitive.ObjectID `bson:"_id,omitempty"`
	ID                 string             `json:"id" bson:"id,omitempty"`
	Creator            string             `json:"creator" bson:"creator,omitempty"`
	LoanAmount         string             `json:"loan_amount" bson:"loan_amount,omitempty"`
	Interest           string             `json:"interest" bson:"interest,omitempty"`
	LoanStartTime      string             `json:"loan_start_time" bson:"loan_start_time,omitempty"`
	LoanDuration       string             `json:"loan_duration" bson:"loan_duration,omitempty"`
	IsInterestProRated bool               `json:"pro_rated" bson:"pro_rated,omitempty"`
	PawnId             string             `json:"pawn_id" bson:"pawn_id,omitempty"`
}

const (
	BidCollectionName = "bids"
)

type Bids struct {
	collection *mongo.Collection
}

func NewBids(dabase mongo.Database) *Bids {
	return &Bids{
		collection: dabase.Collection(BidCollectionName),
	}
}

func (b *Bids) InsertOne(data Bid) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := b.collection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	}
	return "", errors.New("Cannot parse mongodb insertedid to objectid")
}

// find bid by id in smart contract
func (b *Bids) FindOne(id string) (*Bid, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	var bid Bid
	if err := b.collection.FindOne(ctx, filter).Decode(&bid); err != nil {
		return nil, err
	}
	return &bid, nil
}
