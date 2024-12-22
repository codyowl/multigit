package main 

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	git_config_list_command = "git config --list"
)


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
	get_git_details()
}