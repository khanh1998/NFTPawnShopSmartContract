package controller

import (
	"context"
	"log"
	"time"

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

func (b *BidNPawnController) InsertBidToPawnWithSession(connection *mongo.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var session mongo.Session
		var err error
		if session, err = connection.StartSession(); err != nil {
			log.Panic(err)
		}
		if err = session.StartTransaction(); err != nil {
			log.Panic(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
			var newBid model.BidWrite
			err := c.Bind(&newBid)
			if err != nil {
				log.Panic(err)
				return err
			}
			_, err = b.bid.InsertOne(nil, newBid)
			if err != nil {
				log.Panic(err)
				return err
			}
			payload := model.PawnUpdate{
				Bid: newBid.ID,
			}
			err = b.pawn.UpdateOneBy(nil, "id", newBid.Pawn, payload)
			if err != nil {
				log.Panic(err)
				return err
			}
			if err = session.CommitTransaction(sc); err != nil {
				log.Panic(err)
				return err
			}
			return nil
		}); err != nil {
			log.Panic(err)
		}
		session.EndSession(ctx)
	}
	return fn
}
