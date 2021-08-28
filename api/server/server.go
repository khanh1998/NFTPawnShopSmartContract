package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/auth"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/config"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/controller"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	router     *gin.Engine
	connection *config.Connection
	database   *mongo.Database
	tokenMaker auth.Maker
	env        *config.Env
}

func NewServer(con *config.Connection, env *config.Env) (*Server, error) {
	database := con.GetDatabase(env.DatabaseName)
	tokenMaker, err := auth.NewPasetoMaker(env.SymmetricKey)
	if err != nil {
		return nil, err
	}
	server := Server{
		connection: con,
		database:   database,
		tokenMaker: tokenMaker,
		env:        env,
	}
	server.setupRouter()
	return &server, nil
}

func (s *Server) setupRouter() {
	// authMiddleware := middleware.NewAuthMiddleware(s.tokenMaker)
	router := gin.Default()
	router.Use(cors.Default())
	// authRouter := router.Group("/").Use(authMiddleware)

	userModel := model.NewUsers(s.database)
	userController := controller.NewUserController(*userModel)
	// authRouter.GET("/users/:address", userController.FindOne)
	router.GET("/users/:address", userController.FindOneByAddress)
	router.POST("/users", userController.InsertOne)

	pawnModel := model.NewPawns(s.database)
	pawnController := controller.NewPawnController(pawnModel)
	router.GET("/users/:address/pawns", pawnController.FindAllByCreatorAddress)
	router.POST("/pawns", pawnController.InsertOne)
	router.GET("/pawns/:id", pawnController.FindOne)

	bidModel := model.NewBids(*s.database)
	bidController := controller.NewBidController(bidModel)
	router.POST("/bids", bidController.InsertOne)
	router.GET("/bids/:id", bidController.FindOne)

	router.POST("/auth", userController.CreateLogin(s.tokenMaker))
	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
