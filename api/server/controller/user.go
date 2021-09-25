package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/auth"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/service"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
	service *service.User
}

func NewUserController(service *service.User) *UserController {
	return &UserController{
		service: service,
	}
}

func (u *UserController) InsertOne(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Panic(err)
	}
	data, err := u.service.InsertOne(user)
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (u *UserController) FindOneByAddress(c *gin.Context) {
	address := c.Param("address")
	user, err := u.service.FindOneByAddress(address)
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
		user, err := u.service.FindOneByUsername(loginData.Username)
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
