package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

type Object struct {
	Oid     string  `json:"oid"`
	Meanra  float64 `json:"meanra"`
	Meandec float64 `json:"meandec"`
}

func root(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Hello, this is the Aladin service")
}

func aladin(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	oid := r.URL.Query().Get("oid")
	if oid == "" {
		fmt.Fprintf(w, "No oid provided")
		return
	}
	fmt.Printf("Getting aladin for %v\n", oid)
	tmplFile := "templates/aladin.html"
	tmpl, err := template.New("aladin.html").ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	object := getObject(oid)
	radec := map[string]string{
		"ra":  strconv.FormatFloat(object.Meanra, 'g', 6, 32),
		"dec": strconv.FormatFloat(object.Meandec, 'g', 6, 32),
	}
	err = tmpl.Execute(w, radec)
	if err != nil {
		panic(err)
	}

}

func getObject(oid string) Object {
	url := fmt.Sprintf("http://localhost:8002/object/%v", oid)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch %s\n", url)
		log.Println(err)
		return Object{}
	}
	defer resp.Body.Close()
	obj := Object{}
	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		panic(err)
	}
	return obj
}

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		fs.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/aladin/", aladin)
	http.HandleFunc("/", root)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", cors(fs)))
	log.SetPrefix("API Server: ")
	log.Print("initialized")
	log.Fatal(http.ListenAndServe(":8003", nil))
}
