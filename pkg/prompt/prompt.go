package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

// prompts the user for commit details.
func PromptForCommitDetails() (commitType, commitScope, commitMessage string, err error) {
	commitType, err = SelectCommitType()
	if err != nil {
		return "", "", "", err
	}

	commitScope, err = EnterCommitScope()
	if err != nil {
		return "", "", "", err
	}

	commitMessage, err = EnterCommitMessage()
	if err != nil {
		return "", "", "", err
	}

	return
}

// allows the user to select a commit type.
func SelectCommitType() (string, error) {
	fmt.Println("Select a commit type:")

	commitTypeOptions := []string{"✨ feat", "🐛 fix", "📄 docs", "💅 style", "🛠️  refactor", "🎯 perf", "🧪 test", "👷 build", "🔃 ci", "🧹 chore", "🔙 revert"}
	selectedIndex := 0

	for {
		for i, option := range commitTypeOptions {
			marker := "  "
			if i == selectedIndex {
				marker = "➔ "
			}
			fmt.Printf("%s%s\n", marker, option)
		}

		err := keyboard.Open()
		if err != nil {
			return "", err
		}
		defer func() {
			_ = keyboard.Close()
		}()

		_, key, err := keyboard.GetKey()
		if err != nil {
			return "", err
		}

		if key == keyboard.KeyArrowDown && selectedIndex < len(commitTypeOptions)-1 {
			selectedIndex++
		} else if key == keyboard.KeyArrowUp && selectedIndex > 0 {
			selectedIndex--
		} else if key == keyboard.KeyEnter {
			return commitTypeOptions[selectedIndex], nil
		}

		fmt.Print("\033[H\033[2J")
		fmt.Println("Select a commit type:")
	}
}

// allows the user to enter a commit scope.
func EnterCommitScope() (string, error) {
	fmt.Print("Enter the commit scope (optional): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commitScope := strings.TrimSpace(scanner.Text())
	return commitScope, scanner.Err()
}

// allows the user to enter a commit message.
func EnterCommitMessage() (string, error) {
	fmt.Print("Enter the commit message: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commitMessage := strings.TrimSpace(scanner.Text())
	return commitMessage, scanner.Err()
}
