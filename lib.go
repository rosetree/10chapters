package main

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type book struct {
	Name     string
	Chapters int
}

type negDayErr int

func (e negDayErr) Error() string {
	return fmt.Sprintf("cannot create a list for negative day %d", e)
}

// decidePrintDay uses all day parameters to decide which days
// chapters will be printed.
//
// It preferes a positive dayNr over everything else. If
// none is given, a day from the date and its adjustments is
// calculated.
//
// Errors can be anything from the conversion of date to time. Or a
// negDayErr when the adjustments to date result in a negative day.
func decidePrintDay(dayNr int, date string, advanced, skipped int) (day int, err error) {
	day = dayNr
	if day > 1 {
		return
	}

	day, err = daysSince(date)
	if err != nil {
		return
	}

	day = day + advanced - skipped
	if day < 1 {
		err = negDayErr(day)
		return
	}

	return
}

func daysSince(dateStarted string) (int, error) {
	startDate, err := time.Parse("2006-01-02", dateStarted)
	if err != nil {
		return -1, err
	}

	dur := time.Since(startDate)
	// Add 1, because dateStarted should be inclusive (day 1, not 0)
	days := int(dur.Hours()/24) + 1

	return days, nil
}

// printChapters writes the chapters for dayNr to w.
func printChapters(w io.Writer, dayNr int, chapters [10][]string) {
	fmt.Fprintf(w, "Your 10 Chapters for today (day %d):\n", dayNr)
	for listNumber, chapters := range chapters {
		index := (dayNr - 1) % len(chapters)
		chapter := chapters[index]
		fmt.Fprintf(w, "List %d: %s (%d/%d)\n", listNumber, chapter, index+1, len(chapters))
	}
}

type chapterData struct {
	Nr, NrInList, ChaptersInList int
	Chapter                      string
}

type tmplData struct {
	DayNr    int
	Chapters [10]chapterData
}

func prepareTmplData(dayNr int, chapters [10][]string) tmplData {
	var tmplChapters [10]chapterData

	for listNumber, chapters := range chapters {
		index := (dayNr - 1) % len(chapters)
		chapter := chapters[index]

		tmplChapters[listNumber] = chapterData{
			Nr:             listNumber,
			Chapter:        chapter,
			NrInList:       index + 1,
			ChaptersInList: len(chapters),
		}
	}

	return tmplData{dayNr, tmplChapters}
}

func generateListChapters(lists [10][]book) (chapters [10][]string) {
	for listNumber, books := range lists {
		for _, book := range books {
			for chapter := 1; chapter <= book.Chapters; chapter++ {
				name := book.Name + " " + strconv.Itoa(chapter)
				chapters[listNumber] = append(chapters[listNumber], name)
			}
		}
	}
	return chapters
}

func generateLists() (lists [10][]book) {
	lists[0] = []book{
		book{"Matthäus", 28},
		book{"Markus", 16},
		book{"Lukas", 24},
		book{"Johannes", 21},
	}

	lists[1] = []book{
		book{"Genesis", 50},
		book{"Exodus", 40},
		book{"Levitikus", 27},
		book{"Numeri", 36},
		book{"Deuteronomium", 34},
	}

	lists[2] = []book{
		book{"Römer", 16},
		book{"1. Korinther", 16},
		book{"2. Korinther", 13},
		book{"Galater", 6},
		book{"Epheser", 6},
		book{"Philipper", 4},
		book{"Kolosser", 4},
		book{"Hebräer", 13},
	}

	lists[3] = []book{
		book{"1. Thessalonicher", 5},
		book{"2. Thessalonicher", 3},
		book{"1. Timotheus", 6},
		book{"2. Timotheus", 4},
		book{"Titus", 3},
		book{"Philemon", 1},
		book{"Jakobus", 5},
		book{"1. Petrus", 5},
		book{"2. Petrus", 3},
		book{"1. Johannes", 5},
		book{"2. Johannes", 1},
		book{"3. Johannes", 1},
		book{"Judas", 1},
		book{"Offenbarung", 22},
	}

	lists[4] = []book{
		book{"Hiob", 42},
		book{"Prediger", 12},
		book{"Hoheslied", 8},
	}

	lists[5] = []book{
		book{"Psalm", 150},
	}

	lists[6] = []book{
		book{"Sprüche", 31},
	}

	lists[7] = []book{
		book{"Josua", 24},
		book{"Richter", 21},
		book{"Ruth", 4},
		book{"1. Samuel", 31},
		book{"2. Samuel", 24},
		book{"1. Könige", 22},
		book{"2. Könige", 25},
		book{"1. Chronik", 29},
		book{"2. Chronik", 36},
		book{"Esra", 10},
		book{"Nehemia", 13},
		book{"Esther", 10},
	}

	lists[8] = []book{
		book{"Jesaja", 66},
		book{"Jeremia", 52},
		book{"Klagelieder", 5},
		book{"Hesekiel", 48},
		book{"Daniel", 12},
		book{"Hosea", 14},
		book{"Joel", 4},
		book{"Amos", 9},
		book{"Obadja", 1},
		book{"Jona", 4},
		book{"Micha", 7},
		book{"Nahum", 3},
		book{"Habakuk", 3},
		book{"Zefania", 3},
		book{"Haggai", 2},
		book{"Sacharja", 14},
		book{"Maleachi", 3},
	}

	lists[9] = []book{
		book{"Apostelgeschichte", 28},
	}

	return lists
}
