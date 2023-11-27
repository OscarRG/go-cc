package prompt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/OscarRG/go-cc/utils"
	"github.com/eiannone/keyboard"
)

var commitTypeOptions []string

func init() {
	options, err := readCommitTypeOptions()

	if err != nil {
		// use default options
		options = []string{"✨ feat", "🐛 fix", "💡 improvement", "📄 docs", "💅 style", "🛠️  refactor", "🎯 perf", "🧪 test", "👷 build", "🔃 ci", "🧹 chore", "🔙 revert"}
	}

	commitTypeOptions = options
}

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

	selectedIndex := 0

	err := keyboard.Open()
	if err != nil {
		return "", err
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		for i, option := range commitTypeOptions {
			marker := "  "
			if i == selectedIndex {
				marker = "➔ "
			}
			fmt.Printf("%s%s\n", marker, option)
		}

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
		} else if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
			return "", fmt.Errorf("selection canceled")
		}

		fmt.Print("\033[H\033[2J")
		fmt.Println("Select a commit type:")
	}
}

// allows the user to enter a commit scope
func EnterCommitScope() (string, error) {
	fmt.Print("Enter the commit scope (optional): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commitScope := strings.TrimSpace(scanner.Text())
	return commitScope, scanner.Err()
}

// EnterCommitMessage allows the user to enter a non-empty commit message
func EnterCommitMessage() (string, error) {
	for {
		fmt.Print("Enter the commit message: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		commitMessage := strings.TrimSpace(scanner.Text())

		if err := scanner.Err(); err != nil {
			return "", err
		}

		if commitMessage != "" {
			return commitMessage, nil
		}

		fmt.Println("Commit message cannot be empty. Please try again.")
	}
}

// read commit type options from config file
func readCommitTypeOptions() ([]string, error) {
	homeDir := utils.GetHomeDir()
	configFilePath := filepath.Join(homeDir, ".config", "go-cc", "config")

	rootDir, err := utils.GetGitRootDir()
	if err == nil {
		goccFilePath := filepath.Join(rootDir, ".gocc")
		if _, err := os.Stat(goccFilePath); err == nil {
			configFilePath = goccFilePath
		}
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var options []string

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&options); err != nil {
		return nil, fmt.Errorf("error parsing content from %s: %v", configFilePath, err)
	}

	return options, nil
}
