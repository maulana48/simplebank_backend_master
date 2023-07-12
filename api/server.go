package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	db "github.com/maulana48/backend_master_class/simplebank/db/sqlc"
	"github.com/maulana48/backend_master_class/simplebank/token"
	util "github.com/maulana48/backend_master_class/simplebank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	Config     util.Config
	Store      db.Store
	TokenMaker token.Maker
	Router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w", err)
	}

	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

	server.Router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
