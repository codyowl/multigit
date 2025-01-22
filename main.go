package main 

import (
	"fmt"
	"multigit/utilities"
)



func main(){
	choice := utilities.SimpleMenu()

	switch choice{
	case "list":
		utilities.GetProfilesList()
	case "active":
		utilities.GetActiveProfileDetail()
	case "add":
		utilities.AddProfile()
	case "remove":
		utilities.RemoveProfile()
	default:
		fmt.Println("Invalid option !")
		}			
}