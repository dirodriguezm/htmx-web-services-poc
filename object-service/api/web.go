package api

import (
	"alerce/object/core"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

func webRaHms(enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		ra := r.URL.Query().Get("ra")
		tmplFile := "templates/ra_hms.html"
		tmpl, err := template.New("ra_hms.html").ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		raFloat, err := strconv.ParseFloat(ra, 64)
		if err != nil {
			panic(err)
		}
		hms := core.DegreeToHms(raFloat)
		err = tmpl.Execute(w, hms)
		if err != nil {
			panic(err)
		}
	}
}

func webRaDegree(enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		ra := r.URL.Query().Get("ra")
		tmplFile := "templates/ra_deg.html"
		tmpl, err := template.New("ra_deg.html").ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		raFloat, err := core.HmsToDegree(ra)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, raFloat)
		if err != nil {
			panic(err)
		}
	}
}

func webDecDms(enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		dec := r.URL.Query().Get("dec")
		tmplFile := "templates/dec_dms.html"
		tmpl, err := template.New("dec_dms.html").ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		decFloat, err := strconv.ParseFloat(dec, 64)
		if err != nil {
			panic(err)
		}
		dms := core.DegreeToDms(decFloat)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, dms)
		if err != nil {
			panic(err)
		}
	}
}

func webDecDegree(enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		dec := r.URL.Query().Get("dec")
		tmplFile := "templates/dec_deg.html"
		tmpl, err := template.New("dec_deg.html").ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		deg, err := core.DmsToDegrees(dec)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, deg)
		if err != nil {
			panic(err)
		}
	}
}

func webMjdToGreg(enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		mjd := r.URL.Query().Get("mjd")
		tmplFile := "templates/mjd_greg.html"
		tmpl, err := template.New("mjd_greg.html").ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		mjdFloat, err := strconv.ParseFloat(mjd, 64)
		if err != nil {
			panic(err)
		}
		gregDate := core.MjdToGreg(mjdFloat)
		gregDateSplit := strings.Split(gregDate, " ")
		gregDateMap := map[string]string{
			"date": gregDateSplit[0],
			"time": gregDateSplit[1],
		}
		err = tmpl.Execute(w, gregDateMap)
		if err != nil {
			panic(err)
		}
	}
}

func webGregToMjd(enableCors func(w *http.ResponseWriter)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		date := r.URL.Query().Get("date")
		time := r.URL.Query().Get("time")
		greg := date + " " + time
		tmplFile := "templates/greg_mjd.html"
		tmpl, err := template.New("greg_mjd.html").ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		mjd, err := core.GregToMjd(greg)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, mjd)
		if err != nil {
			panic(err)
		}
	}
}
