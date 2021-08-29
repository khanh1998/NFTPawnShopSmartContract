package model

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidWrite struct {
	UUID               primitive.ObjectID `bson:"_id,omitempty"`
	ID                 string             `json:"id" bson:"id,omitempty"`
	Creator            string             `json:"creator" bson:"creator,omitempty"`
	LoanAmount         string             `json:"loan_amount" bson:"loan_amount,omitempty"`
	Interest           string             `json:"interest" bson:"interest,omitempty"`
	LoanStartTime      string             `json:"loan_start_time" bson:"loan_start_time,omitempty"`
	LoanDuration       string             `json:"loan_duration" bson:"loan_duration,omitempty"`
	IsInterestProRated bool               `json:"pro_rated" bson:"pro_rated,omitempty"`
	Pawn               string             `json:"pawn" bson:"pawn,omitempty"`
}

type BidRead BidWrite

func GetBidQueriableParams() []string {
	return []string{
		"id", "creator", "pawn",
	}
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

func (b *Bids) InsertOne(sc mongo.SessionContext, data BidWrite) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result *mongo.InsertOneResult
	var err error
	if sc != nil {
		result, err = b.collection.InsertOne(sc, data)
	} else {
		result, err = b.collection.InsertOne(ctx, data)
	}
	if err != nil {
		return "", err
	}
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	}
	return "", errors.New("cannot parse mongodb insertedid to objectid")
}

// find bid by id in smart contract
func (b *Bids) FindOne(id string) (*BidRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	var bid BidRead
	if err := b.collection.FindOne(ctx, filter).Decode(&bid); err != nil {
		return nil, err
	}
	return &bid, nil
}

func (b *Bids) Find(filter interface{}) ([]BidRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := b.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bids []BidRead
	if err = cursor.All(ctx, &bids); err != nil {
		return nil, err
	}
	return bids, nil
}
