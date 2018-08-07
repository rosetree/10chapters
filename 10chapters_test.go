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
