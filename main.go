package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := getCurrentDatetime()
	fmt.Printf("Now: %v\n", now)

	fmtNow := formatDateToDayMonthYear(now, false)
	fmt.Printf("Formatted now: %v\n", fmtNow)

	fmtNowWithHour := formatDateToDayMonthYear(now, true)
	fmt.Printf("Formatted now with hour: %v\n", fmtNowWithHour)

	birthDate := getSpecificDatetime(2001, 12, 10, 15, 23, 0, false)
	fmt.Printf("Specific date: %v\n", birthDate)

	fmtBirthDate := formatDateToDayMonthYear(birthDate, true)
	fmt.Printf("Formatted specific date: %v\n", fmtBirthDate)

	dateFromStr := getDatetimeFromString("2006-01-02 15:04:05", "2022-10-17 20:29:00")
	fmt.Printf("Date created from string: %v\n", dateFromStr)

	fmtDateFromStr := formatDateToDayMonthYear(dateFromStr, true)
	fmt.Printf("Formatted created from string: %v\n", fmtDateFromStr)

	tomorrow := addToDatetime(now, 1, "days")
	fmt.Printf("Tomorrow will be: %v\n", tomorrow)

	fmtTomorrow := formatDateToDayMonthYear(tomorrow, false)
	fmt.Printf("Formatted tomorrow: %v\n", fmtTomorrow)
}

func getCurrentDatetime() time.Time {
	return time.Now()
}

func getSpecificDatetime(year, month, day, hour, min, sec int, utc bool) time.Time {
	tMonth := time.Month(month)
	loc := time.Local

	if utc {
		loc = time.UTC
	}

	return time.Date(year, tMonth, day, hour, min, sec, 0, loc)
}

func getDatetimeFromString(format, rawDatetime string) time.Time {
	p, err := time.Parse(format, rawDatetime)

	// only for example purposes, this method would return the error
	if err != nil {
		log.Fatal(err)
	}

	return p
}

// formatDateToDayMonthYear formats to d/m/Y
// if hour = true, the format will be d/m/Y H:m:s
func formatDateToDayMonthYear(rawDate time.Time, hour bool) string {
	var hourFormat string

	if hour {
		hourFormat = "15:04:05"
	}

	fmtString := fmt.Sprintf("%s %s", "02/01/2006", hourFormat)

	return rawDate.Format(fmtString)
}

func addToDatetime(datetime time.Time, quant int, what string) time.Time {
	var finalDate time.Time

	if what == "days" {
		daysToAdd := time.Duration(quant) * 24 * time.Hour
		finalDate = datetime.Add(daysToAdd)
	}

	if what == "hours" {
		hoursToAdd := time.Duration(quant) * time.Hour
		finalDate = datetime.Add(hoursToAdd)
	}

	return finalDate
}
