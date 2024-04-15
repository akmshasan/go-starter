package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// const (
// 	DB_SOURCE      = "postgresql://test:secret@localhost:5432/fruit_store?sslmode=disable"
// 	SERVER_ADDRESS = "0.0.0.0:8000"
// )

// func testServer(t *testing.T, store db.Store) *Server {
// 	config := util.Config{
// 		DBSource:      DB_SOURCE,
// 		ServerAddress: SERVER_ADDRESS,
// 	}

// 	server, err := NewServer(config, store)
// 	require.NoError(t, err)

// 	return server
// }

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
