package dbconn

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool() *pgxpool.Pool {
	log.SetPrefix("CreateDatabase: ")
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error acquiring connection: ", err)
	}
	conn.Release()
	return dbpool
}

func CreateTables(pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error acquiring connection: ", err)
	}
	defer conn.Release()
	createObjectTable(conn)
}

func DeleteTables(pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error acquiring connection: ", err)
	}
	defer conn.Release()
	deleteObjectTable(conn)
}

func deleteObjectTable(conn *pgxpool.Conn) {
	stmt := "drop table if exists detection;"
	conn.Exec(context.Background(), stmt)
}

func createObjectTable(conn *pgxpool.Conn) {
	stmt := `create table object (
	oid varchar primary key,
	corrected boolean not null,
	stellar boolean not null,
	ndet integer not null,
	meanra float not null,
	meandec float not null,
	firstmjd float not null,
	lastmjd float not null
);`
	_, err := conn.Exec(context.Background(), stmt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != "42P07" {
				log.Fatal("Could not create object: ", err)
			}
			log.Printf("%s. Continuing initialization", pgErr.Message)
		}
	}
}

func insertObject(conn *pgxpool.Conn) {
	stmt := `insert into object (oid, corrected, stellar, ndet, meanra, meandec, firstmjd, lastmjd) values
('ZTF20aaelulu', false, false, 30, 100, 59000, 59001, 59001.5)
`
	conn.Exec(context.Background(), stmt)
}

func InsertData(conn *pgxpool.Conn) {
	insertObject(conn)
}
