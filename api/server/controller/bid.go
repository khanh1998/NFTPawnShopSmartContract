package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/service"
)

type BidController struct {
	service *service.Bid
}

func NewBidController(service *service.Bid) *BidController {
	return &BidController{
		service: service,
	}
}

func (b *BidController) InsertOne(c *gin.Context) {
	var bidWrite model.BidWrite
	if err := c.BindJSON(&bidWrite); err != nil {
		log.Panic(err)
	}
	bidRead, err := b.service.InsertOne(nil, &bidWrite)
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
	bidRead, err := b.service.UpdateOneById(nil, id, &bidUpdate)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, bidRead)
}

func (b *BidController) FindOne(c *gin.Context) {
	id := c.Param("id")
	bid, err := b.service.FindOneById(nil, id)
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
	bids, err := b.service.FindAllBy(filter)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, bids)
}
