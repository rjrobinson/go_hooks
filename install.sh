#!/bin/bash

set -e

# Determine the user's shell
if [ -n "$BASH_VERSION" ]; then
  shell="bash"
elif [ -n "$ZSH_VERSION" ]; then
  shell="zsh"
else
  echo "Error: unsupported shell. Only Bash and Zsh are supported."
  exit 1
fi

# Build the Go script
go build -o jira-commit-msg jira_ticket.go

# Add the script to the user's PATH
chmod a+x jira-commit-msg
install_path="/usr/local/bin/jira-commit-msg"
mv jira-commit-msg "$install_path"

# Add the script to the user's shell configuration file
if [ "$shell" = "bash" ]; then
  echo "Adding jira-commit-msg to Bash configuration file"
  echo "alias jira-commit-msg='$(command -v jira-commit-msg)'" >> "$HOME/.bashrc"
  source "$HOME/.bashrc"
elif [ "$shell" = "zsh" ]; then
  echo "Adding jira-commit-msg to Zsh configuration file"
  echo "alias jira-commit-msg='$(command -v jira-commit-msg)'" >> "$HOME/.zshrc"
  source "$HOME/.zshrc"
fi

echo "jira-commit-msg installed successfully to $install_path"