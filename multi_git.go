package main 

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"path/filepath"
	"bufio"
)

const (
	mainGitCommand = "git"
)

var gitConfigList = []string{"config", "--list"}

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

// to get documents folder path
func getDocumentsPath() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
	}

	documentsPath := filepath.Join(usr.HomeDir, "Documents")
	return documentsPath
}

// to get profiles list 
func getProfilesList(){
	fmt.Println("Need to read from a notes and list out here")
}

// to get current profile detail
func getActiveProfileDetail(){
	gitConfigListCmd := exec.Command(mainGitCommand, gitConfigList...)
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
			fmt.Println("The active user name is", userName)
			return 
		}
	}
}

func addProfile(){
	fmt.Println("Add a new username and email to the list")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter username: ")
	username, _ := reader.ReadString('\n')
	

	fmt.Println("Enter user.email: ")
	email, _ := reader.ReadString('\n')

	// getting documents home folder path to append the file
	documentsFolder := getDocumentsPath()
	profilesFile := filepath.Join(documentsFolder, "multigit_profiles.txt")
	
	// appending data to a file
	file, err := os.OpenFile(profilesFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()



	// writing the data 
	_, err = file.WriteString(fmt.Sprintf("username=%s, user.email=%s\n", username, email))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Println("Profile added successfully!")
}

func removeProfile(){
	fmt.Println("Remove an username and email from the list")
}

func main(){
	choice := simpleMenu()

	switch choice{
	case "list":
		getProfilesList()
	case "active":
		getActiveProfileDetail()
	case "add":
		addProfile()
	case "remove":
		removeProfile()
	default:
		fmt.Println("Invalid option !")
		}			
}