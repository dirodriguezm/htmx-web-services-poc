package api

import (
	"alerce/lightcurve/dbconn"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func restEndpoints(db *pgxpool.Pool) {
	http.HandleFunc("/", restRootHandler)
	http.HandleFunc("/detections/", restGetDetectionsHandler(db))
}

func webEndpoints(db *pgxpool.Pool) {
	http.HandleFunc("/plot/", webGetLightcurvePlot(db, enableCors))
}

func Api() {
	db := dbconn.CreatePool()
	dbconn.CreateTables(db)
	restEndpoints(db)
	webEndpoints(db)
	log.SetPrefix("API Server: ")
	log.Print("initialized")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
