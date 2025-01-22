package utilities

import (
	"fmt"
	"bufio"
	"path/filepath"
	"os/exec"
	"os"
	"os/user"
	"strings"
)

// git commands
const mainGitCmd = "git"
// for sub commands
var gitConfigListCmd = []string{"config", "--list"}

// to get profiles list 
func GetProfilesList(){
	fmt.Println("Need to read from a notes and list out here")
}

// to get current profile detail
func GetActiveProfileDetail(){
	gitConfigListCmd := exec.Command(mainGitCmd , gitConfigListCmd...)
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

// helper function to add a profile
func AddProfile(){
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

// helper function to get documents folder path
func getDocumentsPath() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
	}

	documentsPath := filepath.Join(usr.HomeDir, "Documents")
	return documentsPath
}

// helper function to remove a profile
func RemoveProfile(){
	fmt.Println("Remove an username and email from the list")
}
