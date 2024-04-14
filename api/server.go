package api

import (
	db "github.com/akmshasan/fruit-store/db/sqlc"
	middleware "github.com/akmshasan/fruit-store/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server serve all HTTP requests for fruit-store
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {

	// Set GIN_MODE=release for production
	gin.SetMode(gin.ReleaseMode)

	server := &Server{store: store}
	router := gin.Default()

	// Add Logger Middleware
	router.Use(middleware.RequestLogger())
	router.Use(middleware.ResponseLogger())

	// Add Prometheus middleware
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Add routes to router
	router.GET("/", server.IndexPage)
	router.GET("/health", server.HealthStatus)
	router.POST("/fruits", server.createFruit)
	router.GET("/fruits/:id", server.getFruit)
	router.GET("/fruits", server.listFruit)
	router.DELETE("/fruits/:id", server.deleteFruit)
	router.PUT("/fruits", server.updateFruit)

	server.router = router
	return server
}

// Runs the server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
