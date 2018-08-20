package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func serve(chapters [10][]string) {
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

		tmpl := template.Must(template.ParseFiles("tmpl/web.html"))
		tmplData := prepareTmplData(day, chapters)
		tmpl.Execute(w, tmplData)
	})

	http.ListenAndServe(":8080", nil)
}
