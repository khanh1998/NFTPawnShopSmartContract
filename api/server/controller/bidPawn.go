package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidNPawnController struct {
	bid  *model.Bids
	pawn *model.Pawns
}

func NewBidNPawnController(bid *model.Bids, pawn *model.Pawns) *BidNPawnController {
	return &BidNPawnController{
		bid:  bid,
		pawn: pawn,
	}
}

func (b *BidNPawnController) InsertBidToPawn(c *gin.Context, sc mongo.SessionContext) error {
	var newBid model.BidWrite
	err := c.Bind(&newBid)
	if err != nil {
		// log.Panic(err)
		return err
	}
	_, err = b.bid.InsertOne(sc, newBid)
	if err != nil {
		// log.Panic(err)
		return err
	}
	payload := model.PawnUpdate{
		Bid: newBid.ID,
	}
	err = b.pawn.UpdateOneBy(sc, "id", newBid.Pawn, payload)
	if err != nil {
		// log.Panic(err)
		return err
	}
	return nil
}

func (b *BidNPawnController) AcceptBid(c *gin.Context, sc mongo.SessionContext) error {
	var bidUpdate model.BidUpdate
	if err := c.Bind(&bidUpdate); err != nil {
		return err
	}
	id := c.Param("id")
	if err := b.bid.UpdateOneBy(sc, "id", id, bidUpdate); err != nil {
		return err
	}
	bid, err := b.bid.FindOne(id)
	if err != nil {
		return err
	}
	pawnUpdate := model.PawnUpdate{
		Status: model.DEAL,
	}
	if err = b.pawn.UpdateOneBy(sc, "id", bid.Pawn, pawnUpdate); err != nil {
		return err
	}
	return nil
}

func (b *BidNPawnController) CancelBid(c *gin.Context, sc mongo.SessionContext) error {
	bidUpdate := model.BidUpdate{
		Status: model.BID_CANCELLED,
	}
	id := c.Param("id")
	if err := b.bid.UpdateOneBy(sc, "id", id, bidUpdate); err != nil {
		return err
	}
	return nil
}
