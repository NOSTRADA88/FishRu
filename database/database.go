package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func Connection(url string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Can't connect to database: %s", err)
	}
	return conn
}

func CloseConnection(conn *pgx.Conn) {
	if err := conn.Close(context.Background()); err != nil {
		log.Fatalf("Can't close connection to DB: %s", err)
	}
}

func Init() {
	connection := Connection(os.Getenv("DB_URL"))
	defer CloseConnection(connection)

	if err := CreateProductTable(connection); err != nil {
		log.Fatalf("Can't create product table: %s", err)
	}
	if err := CreateOwnerTable(connection); err != nil {
		log.Fatalf("Can't create user table: %s", err)
	}
}
