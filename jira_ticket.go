package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func getProjectKey() string {
	file, err := os.Open(".project_key")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return strings.Split(scanner.Text(), "=")[1]
}

func getCommitMessage(commitFile string) string {
	data, err := ioutil.ReadFile(commitFile)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getBranchName() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(output))
}

func main() {
	projectKey := getProjectKey()
	commitFile := os.Args[1]
	commitMessage := getCommitMessage(commitFile)

	branchName := getBranchName()
	ticketNumberRegex := regexp.MustCompile(`(?i)` + projectKey + `-\d+`)
	ticketNumber := ticketNumberRegex.FindString(branchName)

	if len(ticketNumber) > 0 && !strings.Contains(commitMessage, ticketNumber) {
		newCommitMessage := fmt.Sprintf("%s | %s", ticketNumber, commitMessage)
		err := ioutil.WriteFile(commitFile, []byte(newCommitMessage), 0644)
		if err != nil {
			panic(err)
		}
	}
}
