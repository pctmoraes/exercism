package main

import (
	"fmt"
	"booking/booking"
)

func main() {
	schedule := booking.Schedule("7/25/2019 13:45:00")
	fmt.Println(schedule)
}