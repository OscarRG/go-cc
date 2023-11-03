package main

import (
	"fmt"
	"log"

	"github.com/OscarRG/go-cc/pkg/commit"
	"github.com/OscarRG/go-cc/pkg/prompt"
	"github.com/eiannone/keyboard"
)

func main() {
	commitType, commitScope, commitMessage, err := prompt.PromptForCommitDetails()
	if err != nil {
		log.Fatalf("Error during user input: %v", err)
		return
	}

	commitMessageStr := generateCommitMessage(commitType, commitScope, commitMessage)

	fmt.Printf("Generated commit: %s\n", commitMessageStr)

	if askForConfirmation() {
		if err := commit.GenerateAndCommit(commitMessageStr); err != nil {
			log.Fatalf("Error running 'git commit': %v", err)
		} else {
			fmt.Println("Commit successful.")
		}
	} else {
		fmt.Println("Commit canceled.")
	}
}

func generateCommitMessage(commitType, commitScope, commitMessage string) string {
	commitMessageStr := commitType

	if commitScope != "" {
		commitMessageStr += "(" + commitScope + ")"
	}

	commitMessageStr += ": " + commitMessage

	return commitMessageStr
}

func askForConfirmation() bool {
	fmt.Print("Press Enter to confirm or 'C' to cancel: ")
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatal(err)
	}

	return key != keyboard.KeyCtrlC && key != keyboard.KeyCtrlD && char != 'c' && char != 'C'
}
