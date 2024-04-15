package api

import (
	db "github.com/akmshasan/fruit-store/db/sqlc"
	middleware "github.com/akmshasan/fruit-store/middleware"
	"github.com/akmshasan/fruit-store/util"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server serve all HTTP requests for fruit-store
type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{config: config, store: store}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	// Set GIN_MODE=release for production
	// gin.SetMode(gin.ReleaseMode)

	// Initialize router
	router := gin.Default()

	// Set a static favicon.ico
	router.StaticFile("/favicon.ico", "./favicon.ico")

	// Add Logger Middleware
	router.Use(middleware.RequestLogger())
	router.Use(middleware.ResponseLogger())

	// Add Prometheus middleware
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Add routes to router
	router.GET("/", IndexPage)
	router.GET("/health", HealthStatus)
	router.POST("/fruits", server.createFruit)
	router.GET("/fruits/:id", server.getFruit)
	router.GET("/fruits", server.listFruit)
	router.DELETE("/fruits/:id", server.deleteFruit)
	router.PUT("/fruits", server.updateFruit)

	server.router = router
}

// Runs the server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
