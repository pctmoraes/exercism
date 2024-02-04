package main

import (
	"fmt"
	"needfspeed/needfspeed"
)

func main() {
	speed := 5
	batteryDrain := 2
	car := needfspeed.NewCar(speed, batteryDrain)

	distance := 100
	track := needfspeed.NewTrack(distance)

	fmt.Println(needfspeed.CanFinish(car, track))
	car = needfspeed.Drive(car)
	fmt.Println(car)
}