package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	dateTime, _ := time.Parse("1/02/2006 15:04:05", date)
	return dateTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateNow := time.Now()
	dateTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return dateTime.Before(dateNow)

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hr, _, _ := dateTime.Clock()
	return hr >= 12 && hr < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dateTime, _ := time.Parse("1/2/2006 15:04:05", date)
	year, month, day := dateTime.Date()
	hr, min, _ := dateTime.Clock()
	weekDay := dateTime.Weekday()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", weekDay, month, day, year, hr, min)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
