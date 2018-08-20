package main

import (
	"testing"
	"time"
)

// TestDaysSince tests the supported date format and that an
// unsupported date format returns an error and -1. Test that the
// provided date is inclusive.
func TestDaysSince(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	sinceToday, err := daysSince(today)
	if err != nil {
		t.Log("daysSince with", today, "should not cause this error:")
		t.Error(err)
	}
	if sinceToday != 1 {
		t.Error("daysSince(today) should be 1, was", sinceToday)
	}

	germanDate := time.Now().Format("02.01.2006")
	sinceGer, err := daysSince(germanDate)
	if err == nil {
		t.Error("daysSince with unsupported date format should cause an error")
	}
	if sinceGer != -1 {
		t.Error("daysSince with unknown date format should be -1, was", sinceGer)
	}
}

// TestGenerateListChapters checks the lengths of the gernated chapter
// lists and tests for a few expected values.
func TestGenerateListChapters(t *testing.T) {
	lists := [10][]book{
		[]book{book{"One", 1}, book{"Two", 2}},
		[]book{book{"Two", 2}, book{"One", 1}},
		[]book{book{"None", 0}},
		[]book{book{"abc", 26}},
		[]book{book{"abc", 26}},
		[]book{book{"abc", 26}},
		[]book{book{"abc", 26}},
		[]book{book{"abc", 26}},
		[]book{book{"abc", 26}},
		[]book{book{"abc", 26}},
	}

	chapters := generateListChapters(lists)

	lengths := []int{3, 3, 0, 26, 26, 26, 26, 26, 26, 26}
	for i := 0; i < 10; i++ {
		if len(chapters[i]) == lengths[i] {
			continue
		}
		t.Errorf("Length of list %d doesn’t match, expected %d, was %d",
			i, lengths[i], len(chapters[i]))
	}

	if chapters[0][1] != "Two 1" {
		t.Error("chapters[0][1] should be “Two 1”, was", chapters[0][1])
	}
	if chapters[1][1] != "Two 2" {
		t.Error("chapters[1][1] should be “Two 2”, was", chapters[1][1])
	}
}
