package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"studies.com/api"
	db "studies.com/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbSource      = "postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(&store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
