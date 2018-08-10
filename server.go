package main

import (
	"fmt"
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

		var day int64

		dayParam := r.FormValue("day")
		dateParam := r.FormValue("started")

		if dayParam != "" {
			day, err = strconv.ParseInt(dayParam, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if dayParam == "" && dateParam != "" {
			var daysAdvanced, daysSkipped int64

			advancedParam := r.FormValue("advanced")
			skippedParam := r.FormValue("skipped")

			currentDay, err := daysSince(dateParam)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// TODO: Handle these errors
			daysAdvanced, _ = strconv.ParseInt(advancedParam, 10, 64)
			daysSkipped, _ = strconv.ParseInt(skippedParam, 10, 64)

			day = int64(currentDay) + daysAdvanced - daysSkipped
			if day < 1 {
				http.Error(w, "Error: Cannot create lists for a negative day\n", http.StatusBadRequest)
				return
			}
		}

		w.Header().Set("Content-Type", "text/plain")

		fmt.Fprintf(w, "Your 10 Chapters for today (day %d):\n", day)
		for listNumber, chapters := range chapters {
			index := (int(day) - 1) % len(chapters)
			chapter := chapters[index]
			fmt.Fprintf(w, "List %d: %s (%d/%d)\n", listNumber, chapter, index+1, len(chapters))
		}
	})

	http.ListenAndServe(":8080", nil)
}
