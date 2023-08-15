package core

import (
	"alerce/lightcurve/dbconn"
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func mainSetup() {
	os.Setenv("DATABASE_URL", "user=postgres password=postgres host=localhost port=5432 dbname=postgres")
	db = dbconn.CreatePool()
	dbconn.CreateTables(db)
}

func mainTeardown() {
	defer db.Close()
	dbconn.DeleteTables(db)
}

func testSetup(t *testing.T) {
	fmt.Println("Test Setup")
	conn, err := db.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error acquiring connection: ", err)
	}
	dbconn.InsertData(conn)
	t.Cleanup(func() {
		conn.Exec(context.Background(), "truncate table detection;")
		conn.Release()
	})
}

func TestMain(m *testing.M) {
	mainSetup()
	code := m.Run()
	mainTeardown()
	os.Exit(code)
}

func TestGetDetections(t *testing.T) {
	testSetup(t)
	handleError := func(err error) {
		t.Fatal(err)
	}
	handleSuccess := func(result []Detection) {
		if len(result) != 2 {
			t.Fatalf("Detection list length is not correct %v != 2", len(result))
		}
	}
	GetDetections("ZTF20aaelulu", handleSuccess, handleError, db)
}
