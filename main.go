package main

import (
	"context"
	"log"

	api "github.com/akmshasan/fruit-store/api"
	db "github.com/akmshasan/fruit-store/db/sqlc"
	"github.com/akmshasan/fruit-store/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

// var interruptSignals = []os.Signal{
// 	os.Interrupt,
// 	syscall.SIGTERM,
// 	syscall.SIGINT,
// }

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations", err)
	}

	// ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	// defer stop()

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	runDBMigration(config.MigrationURL, config.DBSource)
	store := db.NewStore(connPool)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot run server", err)
	}

}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot migrate db", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up", err)
	}
	log.Println("db migrated successfully")
}
