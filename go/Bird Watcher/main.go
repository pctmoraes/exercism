package main

import (
	"fmt"
	"birdwatcher/birdwatcher"
)

func main() {
	// birdsPerDay := []int{1, 1, 1, 1, 1, 1, 1, 9, 9, 9, 9, 9, 9, 9, 1,2,3,4,5,6,7}
	// fmt.Println(birdwatcher.BirdsInWeek(birdsPerDay, 3))
	// // => 12
	birdsPerDay := []int{2, 5, 0, 7, 4, 1}
	// birdsPerDay := []int{0,1,3}
	fmt.Println(birdwatcher.FixBirdCountLog(birdsPerDay))
	// => [3 5 1 7 5 1]
	
	
}