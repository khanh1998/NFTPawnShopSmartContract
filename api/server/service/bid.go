package service

import (
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bid struct {
	model *model.Bids
}

func NewBid(model *model.Bids) *Bid {
	return &Bid{
		model: model,
	}
}

func (b *Bid) InsertOne(sc mongo.SessionContext, bidWrite *model.BidWrite) (*model.BidRead, error) {
	_, err := b.model.InsertOne(sc, bidWrite)
	if err != nil {
		return nil, err
	}
	bidRead, err := b.model.FindOneBy("id", bidWrite.ID)
	return bidRead, err
}

func (b *Bid) UpdateOneById(sc mongo.SessionContext, id string, bidUpdate *model.BidUpdate) (*model.BidRead, error) {
	err := b.model.UpdateOneBy(sc, "id", id, bidUpdate)
	if err != nil {
		return nil, err
	}
	bidRead, err := b.model.FindOneBy("id", id)
	return bidRead, err
}

func (b *Bid) FindOneById(id string) (*model.BidRead, error) {
	return b.model.FindOneBy("id", id)
}

func (b *Bid) FindAllBy(filter interface{}) ([]model.BidRead, error) {
	return b.model.Find(filter)
}
