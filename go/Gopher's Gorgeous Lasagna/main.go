package main

import (
	"fmt"
	"lasagna/lasagna"
)

func main() {
	var currentTimeInOven int = 20

	fmt.Println(lasagna.RemainingOvenTime(currentTimeInOven))

	var myLayers int = 2

	fmt.Println(lasagna.PreparationTime(myLayers))

	myLayers = 3

	fmt.Println(lasagna.ElapsedTime(myLayers,currentTimeInOven))

}