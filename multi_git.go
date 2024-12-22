package main 

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/nexidian/gocliselect"
	"golang.org/x/sys/windows"
)

const (
	git_config_list_command = "git config --list"
)


// main menu 
func simpleMenu() string {
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




func get_git_details(){
	gitConfigListCmd := exec.Command("git","config","--list")
	stdout, err := gitConfigListCmd.Output()

	if err != nil{
		fmt.Println(err.Error())
		return 
		}
	
	output := string(stdout)
	// splitting the output lines 
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "user.name="){
			userName := strings.TrimPrefix(line, "user.name=")
			fmt.Println("The user name is", userName)
			return 
		}
	}
}

func main(){
	choice := simpleMenu()
	fmt.Printf("Choice: %s\n", choice)
	get_git_details()
}