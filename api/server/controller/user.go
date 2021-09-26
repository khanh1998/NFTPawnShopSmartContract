package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/auth"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/config"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/service"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
	service *service.User
	redis   *config.RedisClient
}

func NewUserController(service *service.User, redis *config.RedisClient) *UserController {
	return &UserController{
		service: service,
		redis:   redis,
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
	var user *model.User = &model.User{}
	var err error
	address := c.Param("address")
	redisKey := fmt.Sprintf("%v/%v", "user", address)
	cacheHit, err := u.redis.Get(redisKey, user)
	if err != nil {
		log.Panic(err)
	}
	if !cacheHit {
		log.Println("cache miss", redisKey)
		user, err = u.service.FindOneByAddress(address)
		u.redis.Put(redisKey, user)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Println("cache hit", redisKey)
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
