package main

import (
	"fmt"
	"purchase/purchase"
)

func main() {
	needLicense := purchase.NeedsLicense("car")
	fmt.Println(needLicense)
	// => true

	needLicense = purchase.NeedsLicense("bike")
	fmt.Println(needLicense)
	// => false

	needLicense = purchase.NeedsLicense("truck")
	fmt.Println(needLicense)
	// => true

	vehicle := purchase.ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
	fmt.Println(vehicle)
	// => "Toyota Corolla is clearly the better choice."

	vehicle = purchase.ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
	fmt.Println(vehicle)
	// => "Volkswagen Beetle is clearly the better choice."

	resell := purchase.CalculateResellPrice(1000, 1)
	fmt.Println(resell)
	// => 800

	resell = purchase.CalculateResellPrice(1000, 5)
	fmt.Println(resell)
	// => 700

	resell = purchase.CalculateResellPrice(1000, 15)
	fmt.Println(resell)
	// => 500

}