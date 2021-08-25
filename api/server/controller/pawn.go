package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
)

type PawnController struct {
	model *model.Pawns
}

func NewPawnController(model *model.Pawns) *PawnController {
	return &PawnController{
		model: model,
	}
}

func (p *PawnController) InsertOne(c *gin.Context) {
	var pawnWrite model.PawnWrite
	if err := c.BindJSON(&pawnWrite); err != nil {
		log.Panic(err)
	}
	pawnWrite.Bids = []string{}
	log.Println(pawnWrite)
	_, err := p.model.InsertOne(pawnWrite)
	if err != nil {
		log.Panic(err)
	}
	pawnRead, err := p.model.FindOne(pawnWrite.ID)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, pawnRead)
}

// find pawn by id in smart contract
func (p *PawnController) FindOne(c *gin.Context) {
	id := c.Param("id")
	pawn, err := p.model.FindOne(id)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, pawn)
}

func (p *PawnController) FindAllByCreatorAddress(c *gin.Context) {
	address := c.Param("address")
	pawns, err := p.model.FindAllByCreatorAddress(address)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, pawns)
}
