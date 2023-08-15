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
	createDetectionsTable(conn)
	createDetectionsIndex(conn)
}

func DeleteTables(pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error acquiring connection: ", err)
	}
	defer conn.Release()
	deleteDetectionsTable(conn)
}

func deleteDetectionsTable(conn *pgxpool.Conn) {
	stmt := "drop table if exists detection;"
	conn.Exec(context.Background(), stmt)
}

func createDetectionsTable(conn *pgxpool.Conn) {
	stmt := `create table detection (
	candid varchar primary key,
	oid varchar not null,
	mjd float not null,
	mag float not null,
	e_mag float not null
);`
	_, err := conn.Exec(context.Background(), stmt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != "42P07" {
				log.Fatal("Could not create detection: ", err)
			}
			log.Printf("%s. Continuing initialization", pgErr.Message)
		}
	}
}

func createDetectionsIndex(conn *pgxpool.Conn) {
	stmt := "CREATE INDEX oid_idx ON detection (oid);"
	_, err := conn.Exec(context.Background(), stmt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != "42P07" {
				log.Fatal("Could not create index oid on table detection: ", err)
			}
			log.Printf("%s. Continuing initialization", err)
		}
	}
}

func createNonDetectionsTable(conn *pgxpool.Conn) {
	stmt := `create table non_detection (
	oid varchar primary key,
	mjd float not null,
	diffmaglim float not null,
	fid int not null
);`
	_, err := conn.Exec(context.Background(), stmt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != "42P07" {
				log.Fatal("Could not create non_detection: ", err)
			}
			log.Printf("%s. Continuing initialization", pgErr.Message)
		}
	}
}

func createNonDetectionsIndex(conn *pgxpool.Conn) {
	stmt := "CREATE INDEX oid_idx ON non_detection (oid);"
	_, err := conn.Exec(context.Background(), stmt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != "42P07" {
				log.Fatal("Could not create index oid on table non_detection: ", err)
			}
			log.Printf("%s. Continuing initialization", err)
		}
	}
}

func insertDetection(conn *pgxpool.Conn) {
	stmt := `insert into detection (candid, oid, mjd, mag, e_mag) values
('lulu1', 'ZTF20aaelulu', 59000, 20.4, 0.5),
('lulu2', 'ZTF20aaelulu', 59001, 19.6, 0.4),
('otro1', 'ZTF23otro', 60000, 15, 0.1),
('otro2', 'ZTF23otro', 60000.5, 16, 0.2);
`
	conn.Exec(context.Background(), stmt)
}

func insertNonDetection(conn *pgxpool.Conn) {
	stmt := `insert into non_detection (oid, mjd, diffmaglim, fid) values
('ZTF20aaelulu', 59000.5, 20.5, 1),
('ZTF20aaelulu', 59001.5, 20.5, 1),
('ZTF23otro', 60000.5, 20.5, 1);
`
	conn.Exec(context.Background(), stmt)
}

func InsertData(conn *pgxpool.Conn) {
	insertDetection(conn)
	insertNonDetection(conn)
}
