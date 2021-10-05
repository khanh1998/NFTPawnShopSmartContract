package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidNPawnController struct {
	bid  *service.Bid
	pawn *service.Pawn
}

func NewBidNPawnController(bid *service.Bid, pawn *service.Pawn) *BidNPawnController {
	return &BidNPawnController{
		bid:  bid,
		pawn: pawn,
	}
}

func (b *BidNPawnController) InsertBidToPawn(c *gin.Context, sc mongo.SessionContext) error {
	var newBid model.BidWrite
	err := c.Bind(&newBid)
	if err != nil {
		return err
	}
	_, err = b.bid.InsertOne(sc, &newBid)
	if err != nil {
		return err
	}
	payload := model.PawnUpdate{
		Bid: newBid.ID,
	}
	_, err = b.pawn.UpdateOneById(sc, newBid.Pawn, &payload)
	if err != nil {
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
	if _, err := b.bid.UpdateOneById(sc, id, &bidUpdate); err != nil {
		return err
	}
	bid, err := b.bid.FindOneById(sc, id)
	if err != nil {
		return err
	}
	pawnUpdate := model.PawnUpdate{
		Status: model.DEAL,
	}
	if _, err = b.pawn.UpdateOneById(sc, bid.Pawn, &pawnUpdate); err != nil {
		return err
	}
	return nil
}

func (b *BidNPawnController) CancelBid(c *gin.Context, sc mongo.SessionContext) error {
	bidUpdate := model.BidUpdate{
		Status: model.BID_CANCELLED,
	}
	id := c.Param("id")
	if _, err := b.bid.UpdateOneById(sc, id, &bidUpdate); err != nil {
		return err
	}
	return nil
}
