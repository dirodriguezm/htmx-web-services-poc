package api

import (
	"alerce/object/core"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func webGetObjectDetails(db *pgxpool.Pool, enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		oid := strings.TrimPrefix(r.URL.Path, "/object-details/")
		tmplFile := "templates/object_details.html"
		tmpl, err := template.New("object_details.html").ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		handleSuccess := func(result core.Object) {
			err = tmpl.Execute(w, result)
			if err != nil {
				panic(err)
			}
		}
		handleError := func(err error) {
			log.Fatal(err)
		}
		core.GetObjectById(oid, handleSuccess, handleError, db)
	}
}
