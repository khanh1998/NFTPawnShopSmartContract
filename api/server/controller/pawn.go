package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/service"
)

type PawnController struct {
	service *service.Pawn
}

func NewPawnController(service *service.Pawn) *PawnController {
	return &PawnController{
		service: service,
	}
}

func (p *PawnController) InsertOne(c *gin.Context) {
	var pawnWrite model.PawnWrite
	if err := c.BindJSON(&pawnWrite); err != nil {
		log.Panic(err)
	}
	pawnWrite.Bids = []string{}
	log.Println(pawnWrite)
	_, err := p.service.InsertOne(pawnWrite)
	if err != nil {
		log.Panic(err)
	}
	pawnRead, err := p.service.FindOne(pawnWrite.ID)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, pawnRead)
}

// find pawn by id in smart contract
func (p *PawnController) FindOne(c *gin.Context) {
	id := c.Param("id")
	pawn, err := p.service.FindOne(id)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, pawn)
}

func (p *PawnController) FindAllByCreatorAddress(c *gin.Context) {
	address := c.Param("address")
	pawns, err := p.service.FindAllByCreatorAddress(address)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, pawns)
}

func (p *PawnController) FindAllBy(c *gin.Context) {
	query := c.Request.URL.Query()
	filter, err := BuildFilterFromGinQuery(query, model.GetPawnQueriableParams())
	if err != nil {
		log.Panic(err)
	}
	pawns, err := p.service.Find(filter)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, pawns)
}

// update by pawn id in smart contract, not UUID in database
func (p *PawnController) UpdateById(c *gin.Context) {
	var data model.PawnUpdate
	if err := c.BindJSON(&data); err != nil {
		log.Panic(err)
	}
	log.Println(data)
	id := c.Param("id")
	pawn, err := p.service.UpdateOneById(nil, id, &data)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, pawn)
}
