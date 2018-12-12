package main

import (
	"html/template"
	"net/http"
	"strconv"
	"encoding/json"
)

func serve(chapters [10][]string) {
	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/app.html")
	})
	http.HandleFunc("/app.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/app.js")
	})
	http.HandleFunc("/app.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/app.css")
	})
	http.HandleFunc("/manifest.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/manifest.json")
	})
	http.HandleFunc("/icon-512.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/icon-512.png")
	})
	http.HandleFunc("/icon-192.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/icon-192.png")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dayParam := r.FormValue("day")
		dateParam := r.FormValue("started")
		advancedParam := r.FormValue("advanced")
		skippedParam := r.FormValue("skipped")
		formatParam := r.FormValue("format")

		var dayNr, daysAdvanced, daysSkipped int64

		if dayParam != "" {
			dayNr, err = strconv.ParseInt(dayParam, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// TODO: Handle these ParseInt errors
		daysAdvanced, _ = strconv.ParseInt(advancedParam, 10, 64)
		daysSkipped, _ = strconv.ParseInt(skippedParam, 10, 64)

		day, err := decidePrintDay(int(dayNr),
			dateParam, int(daysAdvanced), int(daysSkipped))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmplData := prepareTmplData(day, chapters)

		if (formatParam == "json") {
			json.NewEncoder(w).Encode(tmplData)
		} else {
			tmpl := template.Must(template.ParseFiles("tmpl/web.html"))
			tmpl.Execute(w, tmplData)
		}
	})

	http.ListenAndServe(":8080", nil)
}
