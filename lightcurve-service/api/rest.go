package api

import (
	"alerce/lightcurve/core"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func restRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func restGetDetectionsHandler(db *pgxpool.Pool, enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		oid := strings.TrimPrefix(r.URL.Path, "/detections/")
		handle_success := func(result []core.Detection) {
			FormatOutput[[]core.Detection](result, w)
		}
		handle_error := func(err error) {
			log.Fatal(err)
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				log.Printf("Code %s", pgErr.Code)
			}
			http.Error(w, "Internal server error", 500)
		}
		core.GetDetections(oid, handle_success, handle_error, db)
	}
}
