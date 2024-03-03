package main

import (
	"fmt"
	"lasagna_master/lasagna_master"
)

func main() {
	layers := []string{"sauce","noodles","sauce","meat","mozzarella","noodles"}
	preparationTime := lasagna_master.PreparationTime(layers, 0) 
	fmt.Println(preparationTime)

	fmt.Println(lasagna_master.Quantities(layers))

}
