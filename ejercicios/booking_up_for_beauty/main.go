package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	formatedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}
	}
	return formatedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	haspassed := time.Now().Compare(Schedule(date)) 
	if haspassed > 0 {
		return false
	}
	if haspassed < 0 {
		return false
	}
	return true
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}


func main() {
    appointmentDate := "7/13/2020 20:32:00"
    appointmentTime := Schedule(appointmentDate)
    fmt.Println("Appointment time:", appointmentTime.Format("2006-01-02 15:04"))
	fmt.Println(HasPassed("October 3, 2019 20:32:00"))

	fmt.Println("validando si la fecha es futura",HasPassed("7/13/2020 20:32:00"))
}