package api

import (
	"alerce/lightcurve/core"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func webGetLightcurvePlot(db *pgxpool.Pool, enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		oid := strings.TrimPrefix(r.URL.Path, "/plot/")
		tmplFile := "templates/plot.html"
		tmpl, err := template.New("plot.html").ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		handleSuccess := func(result []core.Detection) {
			err = tmpl.Execute(w, result)
			if err != nil {
				panic(err)
			}
		}
		handleError := func(err error) {
			log.Fatal(err)
		}
		core.GetDetections(oid, handleSuccess, handleError, db)
	}
}
