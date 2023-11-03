package commit

import (
	"os"
	"os/exec"
)

// generates and commits a conventional commit.
func GenerateAndCommit(commitMessage string) error {
	cmd := exec.Command("git", "commit", "-m", commitMessage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
