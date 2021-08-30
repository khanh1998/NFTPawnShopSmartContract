package model

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidStatus int

const (
	BID_CREATED BidStatus = iota
	BID_CANCELLED
	BID_ACCEPTED
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
	Status             BidStatus          `json:"status" bson:"status"`
}

type BidRead BidWrite

type BidUpdate struct {
	Status BidStatus `json:"status" bson:""`
}

func GetBidQueriableParams() map[string]string {
	return map[string]string{
		"id":      "string",
		"creator": "string",
		"pawn":    "string",
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

// you only can update status of the pawn and,
// add bid to pawn.
// the key should be unique.
func (b *Bids) UpdateOneBy(sc mongo.SessionContext, key string, value string, data BidUpdate) error {
	filter := bson.M{key: value}
	bid := bson.M{
		"$set": bson.M{
			"status": data.Status,
		},
	}
	log.Println(bid)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var response *mongo.UpdateResult
	var err error
	if sc != nil {
		response, err = b.collection.UpdateOne(sc, filter, bid)
	} else {
		response, err = b.collection.UpdateOne(ctx, filter, bid)

	}
	if err != nil {
		return err
	}
	if response.ModifiedCount == 1 {
		return nil
	}
	return errors.New("didn't update anything")
}
