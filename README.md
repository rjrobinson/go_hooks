Git Hook to add PMT Ticket number
This repository includes a Go script and a git hook integrated 
with Lefthook that checks if a commit message contains a JIRA ticket number ( or what ever PMT you use ).
If a JIRA ticket number is not found in the commit message, the script will prepend the JIRA ticket number derived from the branch name to the commit message.

example branch name "rj/feat/BUG-123"

### Setup
Install Go (if not already installed): https://golang.org/doc/install  
Install Lefthook: https://github.com/Arkweid/lefthook#installation


### Running the Go Script Independently

You can run the Go script independently with the following command:

```bash
go run jira_ticket.go [commit message file path]
```

This will output the commit message, prepending the JIRA ticket number if it's not already included in the message.

### Integrating with Lefthook
To use the Go script with Lefthook, follow these steps:

Compile the Go script into an executable:

```bash
go build -o jira-commit-msg jira_ticket.go && chmod a+x jira-commit-msg
```

Move the jira-commit-msg executable to a directory that's in your system's PATH. For example:

```bash

mv jira-commit-msg /usr/local/bin
```

Add the following to your project's lefthook.yml file:

```yml
commit-msg:
  commands:
    validate-commit-msg:
      run: jira-commit-msg {1}
```
Run lefthook install in your project directory to install the git hooks.


Now, whenever you make a commit, Lefthook will run the jira-commit-msg script to validate the commit message.

### Setting the Project Key
This script reads the project key from a `/.project_key` file in your project root directory. The file should contain a line in the format `PROJECT_KEY=BUG`.Create this file and add your project key before running the script.


### Assumptions and Considerations
This script assumes that your branch names follow the pattern something/JIRA-1234-something.
The script only checks if the commit message begins with a JIRA ticket number. If you include the ticket number elsewhere in the message, the script will still prepend the ticket number to the beginning of the message.

### Troubleshooting
If you encounter any issues while setting up or using this script, please check the troubleshooting guide for Lefthook or create an issue in this repository.