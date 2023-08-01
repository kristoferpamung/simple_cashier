package api

import (
	db "cashier_api/db/sqlc"
	"cashier_api/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// add routes to router
	router.POST("/users", server.createUser)

	router.POST("/customers", server.createCustomer)

	router.POST("/products", server.createProduct)

	router.POST("/orders", server.createOrder)
	router.GET("/orders/:username", server.getOrder)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
