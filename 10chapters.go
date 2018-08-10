package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var currentDay, daysAdvanced, daysSkipped int
	var dateStarted string
	var runHttpd bool

	flag.BoolVar(&runHttpd, "httpd", false, "Whether to start http server on port 8080")
	flag.StringVar(&dateStarted, "date-started", time.Now().Format("2006-01-02"),
		"The date you started reading this plan.")
	flag.IntVar(&daysAdvanced, "days-advanced", 0, "Amount of days you read in advance.")
	flag.IntVar(&daysSkipped, "days-skipped", 0, "Amount of days you skipped the reading.")
	flag.IntVar(&currentDay, "day", 1, "Current day you are reading.")
	flag.Parse()

	lists := generateLists()
	chapters := generateListChapters(lists)

	if (runHttpd) {
		serve(chapters)
		return
	}

	currentDay, err := decidePrintDay(currentDay,
		dateStarted, daysAdvanced, daysSkipped)
	if err != nil {
		fmt.Println("Couldn’t select a day to use, due to this error:")
		fmt.Println(err)
		return
	}

	fmt.Printf("Your 10 Chapters for today (day %d):\n", currentDay)
	for listNumber, chapters := range chapters {
		index := (currentDay - 1) % len(chapters)
		chapter := chapters[index]
		fmt.Printf("List %d: %s (%d/%d)\n", listNumber, chapter, index+1, len(chapters))
	}
}
