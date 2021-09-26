package service

import (
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"go.mongodb.org/mongo-driver/bson"
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
	pawnRead, err := p.FindOne(pawnWrite.ID)
	return pawnRead, err
}

// find pawn by id in smart contract, not UUID in database
func (p *Pawn) FindOne(id string) (*model.PawnRead, error) {
	filter := bson.M{
		"id": id,
	}
	pawns, err := p.model.Find(filter)
	if err != nil {
		return nil, err
	}
	return &pawns[0], nil
}

func (p *Pawn) FindAllByCreatorAddress(address string) ([]model.PawnRead, error) {
	filter := bson.M{
		"creator": address,
	}
	pawns, err := p.model.Find(filter)
	if err != nil {
		return nil, err
	}
	return pawns, nil
}

func (p *Pawn) Find(filter interface{}) ([]model.PawnRead, error) {
	return p.model.Find(filter)
}

func (p *Pawn) UpdateById(id string, data model.PawnUpdate) (*model.PawnRead, error) {
	err := p.model.UpdateOneBy(nil, "id", id, data)
	if err != nil {
		return nil, err
	}
	pawn, err := p.FindOne(id)
	return pawn, err
}
