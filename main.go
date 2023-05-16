package main

import (
	"database/sql"
	"jwt-paseto-token-in-golang/api"
	db "jwt-paseto-token-in-golang/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSourceLocal = "postgresql://postgres:root@localhost:5432/backend-master?sslmode=disable"
	dbSource      = "postgresql://postgres:root@172.17.0.2:5432/backend-master?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
