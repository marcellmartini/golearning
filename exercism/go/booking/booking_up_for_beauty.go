package booking

import (
	"fmt"
	"time"
)

// The layout have to follow the consts on src/pkg/time/format.go file.
var dateLayout = struct {
	schedule               string
	hasPassed              string
	isAfternoonAppointment string
	description            string
	descriptionOutput      string
}{
	"1/02/2006 15:04:05",
	"January 2, 2006 15:04:05",
	"Monday, January 2, 2006 15:04:05",
	"1/2/2006 15:04:05",
	"Monday, January 2, 2006, at 15:04",
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return dateParse(dateLayout.schedule, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	schedule := dateParse(dateLayout.hasPassed, date)
	return schedule.Sub(time.Now()) < 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	schedule := dateParse(dateLayout.isAfternoonAppointment, date)
	return schedule.Hour() >= 12 && schedule.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
// A reminder that I can Create any format that I want using time.Format()
func Description(date string) string {
	schedule := dateParse(dateLayout.description, date)
	return fmt.Sprintf("You have an appointment on %s.",
		schedule.Format(dateLayout.descriptionOutput))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

func dateParse(dateLayout, date string) time.Time {
	dateTime, err := time.Parse(dateLayout, date)
	if err != nil {
		panic(err)
	}
	return dateTime
}
