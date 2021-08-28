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

type PawnStatus int

const (
	CREATED PawnStatus = iota
	CANCELLED
	DEAL
	LIQUIDATED
	REPAID
)

type PawnWrite struct {
	ID           string     `json:"id" bson:"id,omitempty" validate:"number"`                         // id of pawn
	Creator      string     `json:"creator" bson:"creator,omitempty" validate:"eth_addr"`             // wallet address of pawn's creator in hex format
	TokenAddress string     `json:"token_address" bson:"token_address,omitempty" validate:"eth_addr"` // address to smart contract that manages token of creator
	TokenId      string     `json:"token_id" bson:"token_id,omitempty"`                               // token id of creator on smart contract
	Status       PawnStatus `json:"status" bson:"status,omitempty"`                                   // status of pawn
	Bids         []string   `json:"bids" bson:"bids"`
}

type PawnRead struct {
	UUID         primitive.ObjectID `bson:"_id"`                                          // id of pawn in our database
	ID           string             `json:"id" bson:"id,omitempty"`                       // id of pawn
	Creator      User               `json:"creator" bson:"creator,omitempty"`             // wallet address of pawn's creator in hex format
	TokenAddress string             `json:"token_address" bson:"token_address,omitempty"` // address to smart contract that manages token of creator
	TokenId      string             `json:"token_id" bson:"token_id,omitempty"`           // token id of creator on smart contract
	Status       PawnStatus         `json:"status" bson:"status,omitempty"`               // status of pawn
	BidID        []Bid              `json:"bids" bson:"bids,omitempty"`                   // data of bid that re
}

type PawnUpdate struct {
	Bid    string     `json:"bid,omitempty" bson:"bid,omitempty"`
	Status PawnStatus `json:"status,omitempty" bson:"status,omitempty"` // status of pawn
}

type Pawns struct {
	collection *mongo.Collection
}

const (
	PawnsCollectionName = "pawns"
)

func NewPawns(database *mongo.Database) *Pawns {
	collection := database.Collection(PawnsCollectionName)
	return &Pawns{
		collection: collection,
	}
}

// insert a new Pawn.
// return UUID of new pawn or error if it has
func (p *Pawns) InsertOne(data PawnWrite) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := p.collection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	}
	return "", errors.New("cannot convert insertedid to primitive object id")
}

// you only can update status of the pawn and,
// add bid to pawn.
// the key should be unique.
func (p *Pawns) UpdateOneBy(key string, value string, data PawnUpdate) error {
	filter := bson.M{key: value}
	pawn := bson.M{
		"$set": bson.M{
			"status": data.Status,
		},
		"$push": bson.M{
			"bids": data.Bid,
		},
	}
	log.Println(pawn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := p.collection.UpdateOne(ctx, filter, pawn)
	if err != nil {
		return err
	}
	if response.ModifiedCount == 1 {
		return nil
	}
	return errors.New("didn't update anything")
}

func (p *Pawns) FindAllByCreatorAddress(address string) ([]PawnRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := []bson.M{
		{
			"$match": bson.M{
				"creator": address,
			},
		},
		{
			"$lookup": bson.M{
				// Define the tags collection for the join.
				"from": BidCollectionName,
				// Specify the variable to use in the pipeline stage.
				"let": bson.M{
					"bids": "$bids",
				},
				"pipeline": []bson.M{
					// Select only the relevant bids from the bids collection.
					// Otherwise all the bids are selected.
					{
						"$match": bson.M{
							"$expr": bson.M{
								"$in": []interface{}{
									"$id",
									"$$bids",
								},
							},
						},
					},
					// Sort bids by their id field in asc. -1 = desc
					{
						"$sort": bson.M{
							"id": 1,
						},
					},
				},
				// Use bids as the field name to match struct field.
				"as": "bids",
			},
		},
		{
			"$lookup": bson.M{
				// Define the tags collection for the join.
				"from": UserCollectionName,
				// Specify the variable to use in the pipeline stage.
				"localField":   "creator",
				"foreignField": "wallet_address",
				// Use bids as the field name to match struct field.
				"as": "creator",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$creator",
				"preserveNullAndEmptyArrays": true,
			},
		},
	}
	curr, err := p.collection.Aggregate(ctx, query)
	if err != nil {
		return nil, err
	}
	var pawn []PawnRead
	if err = curr.All(context.Background(), &pawn); err != nil {
		return nil, err
	}
	return pawn, nil
}

// find pawn by id in smart contract, not UUID in database
func (p *Pawns) FindOne(id string) (*PawnRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := []bson.M{
		{
			"$match": bson.M{
				"id": id,
			},
		},
		{
			"$lookup": bson.M{
				// Define the tags collection for the join.
				"from": BidCollectionName,
				// Specify the variable to use in the pipeline stage.
				"let": bson.M{
					"bids": "$bids",
				},
				"pipeline": []bson.M{
					// Select only the relevant bids from the bids collection.
					// Otherwise all the bids are selected.
					{
						"$match": bson.M{
							"$expr": bson.M{
								"$in": []interface{}{
									"$id",
									"$$bids",
								},
							},
						},
					},
					// Sort bids by their id field in asc. -1 = desc
					{
						"$sort": bson.M{
							"id": 1,
						},
					},
				},
				// Use bids as the field name to match struct field.
				"as": "bids",
			},
		},
		{
			"$lookup": bson.M{
				// Define the tags collection for the join.
				"from": UserCollectionName,
				// Specify the variable to use in the pipeline stage.
				"localField":   "creator",
				"foreignField": "wallet_address",
				// Use bids as the field name to match struct field.
				"as": "creator",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$creator",
				"preserveNullAndEmptyArrays": true,
			},
		},
	}
	curr, err := p.collection.Aggregate(ctx, query)
	if err != nil {
		return nil, err
	}
	var pawn []PawnRead
	if err = curr.All(context.Background(), &pawn); err != nil {
		return nil, err
	}
	return &pawn[0], nil
}

func (p *Pawns) Find(filter interface{}) ([]PawnRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := p.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var pawns []PawnRead
	if err = cursor.All(ctx, &pawns); err != nil {
		return nil, err
	}
	return pawns, nil
}
