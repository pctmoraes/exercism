package partyrobot

import (
	"fmt"
	"strconv"
	"strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	welcomeMsg := fmt.Sprintf("Welcome to my party, %s!", name)
	return welcomeMsg
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	hbMsg := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return hbMsg
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcomeMsg := Welcome(name)

	tableNumberFloat := float64(table) / 100.0
	tableNumberString := strconv.FormatFloat(tableNumberFloat, 'f', -1, 64)
	tableNumber := strings.ReplaceAll(tableNumberString, ".", "")
	
	assignTable := fmt.Sprintf("\nYou have been assigned to table %s. Your table is %s, exactly %.1f meters from here.", tableNumber, direction, distance)
	
	assignMate := fmt.Sprintf("\nYou will be sitting next to %s.", neighbor)

	wholeMsg := welcomeMsg + assignTable + assignMate
	
	return wholeMsg
}
