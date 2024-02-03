package main

import (
	"fmt"
	"cars/cars"
)

func main() {
	// fmt.Println(cars.CalculateWorkingCarsPerHour(1547, 90))
	// fmt.Println(cars.CalculateWorkingCarsPerMinute(1105, 90))


	fmt.Println(cars.CalculateCost(37))
	// => 355000
	
	fmt.Println(cars.CalculateCost(21))
	// => 200000

}