package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	projectKeyFile    = ".project_key"
	skippedBranchList = "main,master,hotfix"
)

var (
	jiraTicketRegex = regexp.MustCompile(`([A-Z]+-\d+)`)
)

func main() {
	branchName, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		fmt.Println("error getting branch name:", err)
		os.Exit(1)
	}

	// Check if the branch is included in skipped branches list
	skippedBranches := strings.Split(skippedBranchList, ",")
	for _, branch := range skippedBranches {
		if strings.Contains(string(branchName), branch) {
			fmt.Println("skipped branch")
			os.Exit(0)
		}
	}

	projectKey := getProjectKey()

	commitMsg := getCommitMessage(os.Args[1])
	if !jiraTicketRegex.MatchString(commitMsg) {
		ticket := jiraTicketRegex.FindString(string(branchName))
		newCommitMsg := fmt.Sprintf("%s %s", projectKey, ticket)
		if len(commitMsg) > 0 {
			newCommitMsg += fmt.Sprintf(" - %s", strings.TrimSpace(commitMsg))
		}
		err = ioutil.WriteFile(os.Args[1], []byte(newCommitMsg), 0644)
		if err != nil {
			fmt.Println("error updating commit message:", err)
			os.Exit(1)
		}
	}
}

func getProjectKey() string {
	projectKey := "VL"
	projectKeyBytes, err := ioutil.ReadFile(projectKeyFile)
	if err == nil {
		projectKey = strings.TrimSpace(string(projectKeyBytes))
		if len(projectKey) == 0 {
			projectKey = "VL"
		}
	}

	// Check if the branch is included in skipped branches list
	skippedBranches := strings.Split(skippedBranchList, ",")
	for _, branch := range skippedBranches {
		if strings.Contains(string(projectKey), branch) {
			return string(projectKeyBytes)
		}
	}

	return projectKey
}

func getCommitMessage(fileName string) string {
	commitMsgBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return ""
	}
	return string(commitMsgBytes)
}
