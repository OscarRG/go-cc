package utils

import (
	"log"
	"os"
	"os/exec"
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
