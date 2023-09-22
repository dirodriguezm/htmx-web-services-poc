package api

import (
	"alerce/object/core"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func restRoot(enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		fmt.Fprintf(w, "Hello, this is the Object service")
	}
}

func restGetObjectHandler(db *pgxpool.Pool, enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		oid := strings.TrimPrefix(r.URL.Path, "/object/")
		handle_success := func(result core.Object) {
			FormatOutput[core.Object](result, w)
		}
		handle_error := func(err error) {
			log.Fatal(err)
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				log.Printf("Code %s", pgErr.Code)
			}
			http.Error(w, "Internal server error", 500)
		}
		core.GetObjectById(oid, handle_success, handle_error, db)
	}
}
