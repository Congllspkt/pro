package dbConfig

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func DbConfig() *pgx.Conn {
	connString := "postgresql://myuser:mypassword@localhost:5432/mydatabase"

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln("Error connect to DB")
	}

	// defer db.Close(context.Background())

	log.Println("Connected to DB")
	return db
}