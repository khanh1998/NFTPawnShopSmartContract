package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/uss-kelvin/NFTPawningShopBackend/server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username      string             `json:"username" bson:"username,omitempty" validate:"required,unique,alphanum"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Email         string             `json:"email" bson:"email,omitempty" validate:"email,unique"`
	Password      string             `json:"password" bson:"password,omitempty"`
	WalletAddress string             `json:"wallet_address" bson:"wallet_address,omitempty" validate:"eth_addr"`
}

type Users struct {
	collection *mongo.Collection
}

const (
	UserCollectionName = "users"
)

func NewUsers(database *mongo.Database) *Users {
	collection := database.Collection(UserCollectionName)
	return &Users{
		collection: collection,
	}
}

func (u *Users) FindOne(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	result := u.collection.FindOne(ctx, filter)
	var user User
	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *Users) FindByUsername(username string) (*User, error) {
	filter := bson.M{"username": username}
	users, err := u.Find(filter)
	if err != nil {
		return nil, err
	}
	return &users[0], nil
}

// key is document property
func (u *Users) FindBy(key string, value string) (*User, error) {
	filter := bson.M{key: value}
	users, err := u.Find(filter)
	if err != nil {
		return nil, err
	}
	return &users[0], nil
}

func (u *Users) Find(filter interface{}) ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := u.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *Users) InsertOne(data User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	hashed, err := utils.Hash(data.Password)
	if err != nil {
		return "", err
	}
	data.Password = hashed
	fmt.Println(data)
	result, err := u.collection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	}
	return "", errors.New("cannot convert insertedid to primitive object id")
}
