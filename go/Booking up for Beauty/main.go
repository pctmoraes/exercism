package main

import (
	"fmt"
	"booking/booking"
)

func main() {
	dateAsString := "7/25/2019 13:45:00"
	schedule := booking.Schedule(dateAsString)
	fmt.Println(schedule)

	fmt.Println(booking.HasPassed(dateAsString))
}