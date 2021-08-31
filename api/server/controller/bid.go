package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
)

type BidController struct {
	model *model.Bids
}

func NewBidController(model *model.Bids) *BidController {
	return &BidController{
		model: model,
	}
}

func (b *BidController) InsertOne(c *gin.Context) {
	var bidWrite model.BidWrite
	if err := c.BindJSON(&bidWrite); err != nil {
		log.Panic(err)
	}
	_, err := b.model.InsertOne(nil, bidWrite)
	if err != nil {
		log.Panic(err)
	}
	bidRead, err := b.model.FindOne(bidWrite.ID)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, bidRead)
}

func (b *BidController) UpdateOne(c *gin.Context) {
	var bidUpdate model.BidUpdate
	if err := c.BindJSON(&bidUpdate); err != nil {
		log.Panic(err)
	}
	id := c.Param("id")
	err := b.model.UpdateOneBy(nil, "id", id, bidUpdate)
	if err != nil {
		log.Panic(err)
	}
	bidRead, err := b.model.FindOne(id)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, bidRead)
}

func (b *BidController) FindOne(c *gin.Context) {
	id := c.Param("id")
	bid, err := b.model.FindOne(id)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, bid)
}

func (b *BidController) FindAllBy(c *gin.Context) {
	query := c.Request.URL.Query()
	filter, err := BuildFilterFromGinQuery(query, model.GetBidQueriableParams())
	if err != nil {
		log.Panic(err)
	}
	bids, err := b.model.Find(filter)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, bids)
}
