package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/OscarRG/go-cc/pkg/commit"
	"github.com/OscarRG/go-cc/pkg/prompt"
	"github.com/OscarRG/go-cc/utils"
	"github.com/eiannone/keyboard"
)

const (
	configFolderPath = ".config/go-cc"
	configFileName   = "config"
	exitSuccess      = 0
)

// main is the entry point of the program. It initializes the configuration, prompts the user for commit details,
// generates a commit message, and commits the changes based on user confirmation.
func main() {

	gitRootDir, err := utils.GetGitRootDir()
	if err != nil {
		log.Fatal("Error: This command must be run from within a Git repository.")
	}

	if !utils.HasStagedChanges(gitRootDir) {
		log.Fatal("No changes staged for commit. Please stage your changes and try again.")
	}

	if err := initializeConfig(); err != nil {
		log.Fatalf("Initialization error: %v", err)
	}

	commitType, commitScope, commitMessage, err := prompt.PromptForCommitDetails()
	if err != nil {
		log.Fatalf("Error during user input: %v", err)
	}

	commitMessageStr := generateCommitMessage(commitType, commitScope, commitMessage)
	fmt.Printf("Generated commit: %s\n", commitMessageStr)

	if askForConfirmation() {
		if err := commitAndLog(commitMessageStr); err != nil {
			log.Fatalf("Error committing: %v", err)
		}
	} else {
		fmt.Println("Commit canceled.")
		os.Exit(exitSuccess)
	}
}

// initializeConfig sets up the configuration directory and default config file if they don't exist
func initializeConfig() error {
	configPath := filepath.Join(utils.GetHomeDir(), configFolderPath)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configPath, 0755); err != nil {
			return fmt.Errorf("unable to create config directory: %v", err)
		}
	}

	filePath := filepath.Join(configPath, configFileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		created, err := createDefaultConfigFile(filePath)
		if err != nil {
			return fmt.Errorf("error creating default config file: %v", err)
		}
		if created {
			fmt.Println(" ✨✨ Default configuration has been added ✨✨")
			fmt.Println(" ✨✨ From now on you will be able to use the default conventional commits ✨✨")
			os.Exit(exitSuccess)
		}
	}
	return nil
}

// commitAndLog commits the changes and logs the result
func commitAndLog(commitMessage string) error {
	if err := commit.GenerateAndCommit(commitMessage); err != nil {
		return err
	}
	fmt.Println("Commit successful.")
	return nil
}

// generateCommitMessage generates a formatted commit message based on user input
func generateCommitMessage(commitType, commitScope, commitMessage string) string {
	commitMessageStr := commitType

	if commitScope != "" {
		commitMessageStr += "(" + commitScope + ")"
	}

	commitMessageStr += ": " + commitMessage

	return commitMessageStr
}

// askForConfirmation prompts the user to confirm the commit
func askForConfirmation() bool {
	fmt.Print("Press Enter to confirm or 'C' to cancel: ")
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatal(err)
	}

	return key != keyboard.KeyCtrlC && key != keyboard.KeyCtrlD && char != 'c' && char != 'C'
}

// createDefaultConfigFile creates the default config file with the specified content
// Returns true if the default configuration was created, false otherwise
func createDefaultConfigFile(filePath string) (bool, error) {
	content := `["✨ feat", "🐛 fix", "💡 improvement" ,"📄 docs", "💅 style", "🛠️  refactor", "🎯 perf", "🧪 test", "👷 build", "🔃 ci", "🧹 chore", "🔙 revert"]`

	file, err := os.Create(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return false, err
	}

	return true, nil
}
