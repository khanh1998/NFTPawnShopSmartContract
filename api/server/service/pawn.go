package service

import (
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pawn struct {
	model *model.Pawns
}

func NewPawn(model *model.Pawns) *Pawn {
	return &Pawn{
		model: model,
	}
}

func (p *Pawn) InsertOne(pawnWrite model.PawnWrite) (*model.PawnRead, error) {
	pawnWrite.Bids = []string{}
	_, err := p.model.InsertOne(pawnWrite)
	if err != nil {
		return nil, err
	}
	pawnRead, err := p.FindOne(nil, pawnWrite.ID)
	return pawnRead, err
}

// find pawn by id in smart contract, not UUID in database
func (p *Pawn) FindOne(sc mongo.SessionContext, id string) (*model.PawnRead, error) {
	filter := bson.M{
		"id": id,
	}
	pawns, err := p.model.Find(sc, filter)
	if err != nil {
		return nil, err
	}
	return &pawns[0], nil
}

func (p *Pawn) FindAllByCreatorAddress(address string) ([]model.PawnRead, error) {
	filter := bson.M{
		"creator": address,
	}
	pawns, err := p.model.Find(nil, filter)
	if err != nil {
		return nil, err
	}
	return pawns, nil
}

func (p *Pawn) Find(filter interface{}) ([]model.PawnRead, error) {
	return p.model.Find(nil, filter)
}

func (p *Pawn) UpdateOneById(sc mongo.SessionContext, id string, data *model.PawnUpdate) (*model.PawnRead, error) {
	err := p.model.UpdateOneBy(sc, "id", id, data)
	if err != nil {
		return nil, err
	}
	pawn, err := p.FindOne(sc, id)
	return pawn, err
}
