package main

import (
	"context"
	"log"

	"alerce/lightcurve/dbconn"
)

func main() {
	db := dbconn.CreatePool()
	defer db.Close()
	dbconn.CreateTables(db)
	log.SetPrefix("Insert Data: ")
	conn, err := db.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error acquiring connection: ", err)
	}
	dbconn.InsertData(conn)
	conn.Release()
}
