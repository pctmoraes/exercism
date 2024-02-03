package main

import (
	"fmt"
	"annalyn/annalyn"
)

func main(){
	// var knightIsAwake = false
	// fmt.Println(annalyn.CanFastAttack(knightIsAwake))

	// var knightIsAwake = true
	// var archerIsAwake = true
	// var prisonerIsAwake = true
	// fmt.Println(annalyn.CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))

	// var archerIsAwake = true
	// var prisonerIsAwake = true
	// fmt.Println(annalyn.CanSignalPrisoner(archerIsAwake, prisonerIsAwake))

	var knightIsAwake = false
	var archerIsAwake = true
	var prisonerIsAwake = true
	var petDogIsPresent = false
	fmt.Println(annalyn.CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
}