package service

import (
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
)

type User struct {
	model *model.Users
}

func NewUserService(model *model.Users) *User {
	return &User{
		model: model,
	}
}

func (u *User) InsertOne(input model.User) (*model.User, error) {
	insertedId, err := u.model.InsertOne(input)
	if err != nil {
		return nil, err
	}
	data, err := u.model.FindOne(insertedId)
	return data, err
}

func (u *User) FindOneByAddress(address string) (*model.User, error) {
	return u.model.FindBy("wallet_address", address)
}

func (u *User) FindOneByUsername(username string) (*model.User, error) {
	return u.model.FindBy("wallet_address", username)
}
