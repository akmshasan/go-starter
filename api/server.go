package api

import (
	db "github.com/akmshasan/fruit-store/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serve all HTTP requests for fruit-store
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Add routes to router
	router.POST("/fruits", server.createFruit)
	router.GET("/fruits/:id", server.getFruit)
	router.GET("/fruits", server.listFruit)
	router.DELETE("/fruits/:id", server.deleteFruit)
	// router.PUT("/fruits/:id", server.updateFruit)

	//
	server.router = router
	return server
}

// Runs the server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
