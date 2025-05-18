package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

type CommitOption struct {
	Key   string
	Value string
	Label string
}

// GetCommitOptions returns the list of commit type options
func GetCommitOptions() []CommitOption {
	return []CommitOption{
		{Key: "feat", Value: "💎 feat:", Label: "💎 Feature"},
		{Key: "add", Value: "🎁 add:", Label: "🎁 Minor improvement"},
		{Key: "update", Value: "🆙 update:", Label: "🆙 Update"},
		{Key: "ref", Value: "🔧 ref:", Label: "🔧 Refactor"},
		{Key: "wip", Value: "⏳ wip:", Label: "⏳ Work In Progress"},
		{Key: "delete", Value: "🔥 delete:", Label: "🔥 Deletion"},
		{Key: "chore", Value: "🧹 chore:", Label: "🧹 Chore"},
		{Key: "bugfix", Value: "🐛 bugfix:", Label: "🐛 Bugfix"},
	}
}

func main() {
	fmt.Println("Select your commit type:")

	commitOptions := GetCommitOptions()
	
	// Create a slice of options for the survey
	options := make([]string, len(commitOptions))
	for i, opt := range commitOptions {
		options[i] = opt.Label
	}

	// Create a variable to store the selected option
	var selectedOption string
	
	// Create the selection prompt
	prompt := &survey.Select{
		Message: "Select your commit type:",
		Options: options,
		Default: options[0],
	}
	
	// Ask the user to select a commit type
	err := survey.AskOne(prompt, &selectedOption)
	if err != nil {
		if err == terminal.InterruptErr {
			fmt.Println("Commit cancelled.")
			os.Exit(0)
		}

		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	// Find the selected commit option
	var selectedCommitOption CommitOption
	for _, opt := range commitOptions {
		if opt.Label == selectedOption {
			selectedCommitOption = opt
			break
		}
	}
	
	// Ask for commit message
	var commitMessage string
	messagePrompt := &survey.Input{
		Message: "Commit message:",
	}
	
	surveyOpts := []survey.AskOpt{
		survey.WithValidator(func(val interface{}) error {
			str, _ := val.(string)
			if strings.TrimSpace(str) == "" {
				return fmt.Errorf("Please enter a commit message")
			}
			return nil
		}),
		survey.WithKeepFilter(true),
		// Add custom key handling to ignore problematic key sequences
		survey.WithFilter(func(filter string, value string, index int) (include bool) {
			return true // Always include options regardless of filter
		}),
	}
	
	err = survey.AskOne(messagePrompt, &commitMessage, surveyOpts...)
	
	if err != nil {
		if err == terminal.InterruptErr {
			fmt.Println("Interrupted")
			os.Exit(1)
		}
		// Handle special case for unexpected escape sequences
		if strings.Contains(err.Error(), "unexpected escape sequence") {
			// Just log and continue - the input is still valid
			fmt.Println("\nNote: Some keyboard shortcuts may not work properly in this terminal")
			// User might have pressed Ctrl+C or another interrupt
			if commitMessage == "" {
				fmt.Println("No commit message provided. Exiting.")
				os.Exit(1)
			}
		} else {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	// Execute git add .
	gitAddCmd := exec.Command("git", "add", ".")
	gitAddErr := gitAddCmd.Run()
	
	if gitAddErr != nil {
		fmt.Println("Something went wrong with git add! 😿")
		os.Exit(1)
	}
	
	// Execute git commit
	fullCommitMessage := fmt.Sprintf("%s %s", selectedCommitOption.Value, commitMessage)
	gitCommitCmd := exec.Command("git", "commit", "-m", fullCommitMessage)
	gitCommitErr := gitCommitCmd.Run()
	
	if gitCommitErr != nil {
		fmt.Println("Something went wrong! 😿")
		os.Exit(1)
	}
	
	fmt.Println("Done! 💪")
}