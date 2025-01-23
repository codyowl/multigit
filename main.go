package main 

import (
	"fmt"
	"multigit/utilities"
)



func main(){
	choice := utilities.SimpleMenu()

	// checking if the cli has been triggered from a git repository
	if utilities.IsGitRepository() {
		fmt.Println("Indeed this is a repository ! we can call the function in the future")

	}else {
		fmt.Println("Please trigger the cli from a valid git repository")
	}

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