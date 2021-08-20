package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/auth"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
	model model.Users
}

func NewUserController(model model.Users) *UserController {
	return &UserController{
		model: model,
	}
}

func (u *UserController) InsertOne(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Panic(err)
	}
	insertedId, err := u.model.InsertOne(user)
	if err != nil {
		log.Panic(err)
	}
	data, err := u.model.FindOne(insertedId)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (u *UserController) FindOne(c *gin.Context) {
	id := c.Param("id")
	user, err := u.model.FindOne(id)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserController) CreateLogin(tokenMaker auth.Maker) func(c *gin.Context) {
	return func(c *gin.Context) {
		var loginData LoginData
		if err := c.BindJSON(&loginData); err != nil {
			log.Panic(err)
		}
		user, err := u.model.FindByUsername(loginData.Username)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(loginData.Password)
		fmt.Println(user.Password)
		err = utils.ComparePassword(user.Password, loginData.Password)
		if err != nil {
			log.Panic(err)
			c.Status(http.StatusUnauthorized)
		}
		token, err := tokenMaker.CreateToken(user.Username, time.Hour)
		if err != nil {
			log.Panic(err)
		}
		c.IndentedJSON(http.StatusOK, bson.M{"token": token})
	}
}
