package api

import (
	"html/template"
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
		err = tmpl.Execute(w, oid)
		if err != nil {
			panic(err)
		}
	}
}
