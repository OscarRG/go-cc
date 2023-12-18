package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetHomeDir returns the home directory of the current user
func GetHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %v", err)
	}
	return home
}

// GetGitRootDir returns the root directory of the Git repository
func GetGitRootDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// HasStagedChanges returns true if there are staged changes in the Git repository
func HasStagedChanges(gitRootDir string) bool {
	_, err := exec.Command("git", "diff", "--cached", "--exit-code").Output()
	return err != nil
}

// read commit type options from config file
func ReadCommitTypeOptions() ([]string, error) {
	homeDir := GetHomeDir()
	configFilePath := filepath.Join(homeDir, ".config", "go-cc", "config")

	rootDir, err := GetGitRootDir()
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
