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
	// http.HandleFunc("/", restRootHandler)
	http.HandleFunc("/detections/", restGetDetectionsHandler(db, enableCors))
}

func webEndpoints(db *pgxpool.Pool) {
	http.HandleFunc("/plot/", webGetLightcurvePlot(db, enableCors))
}

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		fs.ServeHTTP(w, r)
	}
}

func staticFiles() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", cors(fs)))
}

func Api() {
	db := dbconn.CreatePool()
	// dbconn.CreateTables(db)
	restEndpoints(db)
	webEndpoints(db)
	staticFiles()
	log.SetPrefix("API Server: ")
	log.Print("initialized")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
