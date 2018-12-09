package main

// Dates and Times in Go are managed in the Time package
// The Time type encapsulates everything you need for both dates
// and times, including math operations, time zones, and so on

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)
	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v\n",
		tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	// In order to create your own date/time format, you modal it
	// using a specific date/time modal
	// such as -- Monday, January 2nd at 3:04 PM in 2006, Mountain Standard Time
	//      or -- Monday, January 2nd 2006
	// Modal is an example string that we provide to the format function
	// It has to follow the specific format to be parsable by the format function

	// This is the standard pattern needs to be followed while specifying date/time
	/*
		1 - Month
		2 - Day
		3 - Hours
		4 - Minutes
		5 - Seconds
		6 - Year
		7 - Time zone
	*/

	longDateFormat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longDateFormat))

	shortDateFormat := "1/2/06"
	fmt.Println("Tomorrow is", tomorrow.Format(shortDateFormat))

	// 12-hour am-pm clock format
	dateTimeFormat := "2006-01-02 03:04:05 PM"
	fmt.Println(tomorrow.Format(dateTimeFormat))

	// 24-hour military time format
	dateTimeFormat24 := "2006-01-02 15:04:05"
	fmt.Println(tomorrow.Format(dateTimeFormat24))

	// location loading
	loc, _ := time.LoadLocation("Pacific/Auckland")
	aucklandTime := time.Now().In(loc)
	fmt.Println(aucklandTime)
}
