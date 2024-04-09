package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/akmshasan/fruit-store/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

// const dbSource = "postgresql://test:secret@localhost:5432/fruit_store?sslmode=disable"

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../.")
	if err != nil {
		log.Fatal("cannot load configurations", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
