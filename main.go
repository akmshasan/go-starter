package main

import (
	"context"
	"log"

	api "github.com/akmshasan/fruit-store/api"
	db "github.com/akmshasan/fruit-store/db/sqlc"
	"github.com/akmshasan/fruit-store/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}

}
