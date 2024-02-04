package main

import (
	"fmt"
	"techpalace/techpalace"
)

func main() {
	fmt.Println(techpalace.WelcomeMessage("Judy"))

	fmt.Println(techpalace.AddBorder("Welcome!", 10))

	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	
	fmt.Println(techpalace.CleanupMessage(message))

}