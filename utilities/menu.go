package utilities

import (
	"fmt"
	"os"
)

// main menu 
func SimpleMenu() string {
	options := []string{"list", "active", "add", "remove"}
	fmt.Println("What would you like to do?")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(options) {
		fmt.Println("Invalid choice. Exiting.")
		os.Exit(1)
	}
	return options[choice-1]
}