package prompt

import (
	"fmt"

	"github.com/OscarRG/go-cc/utils"
	"github.com/charmbracelet/huh"
)

var commitTypeOptions []string

func init() {
	options, err := utils.ReadCommitTypeOptions()
	if err != nil {
		// use default options
		options = []string{"✨ feat", "🐛 fix", "💡 improvement", "📄 docs", "💅 style", "🛠️  refactor", "🎯 perf", "🧪 test", "👷 build", "🔃 ci", "🧹 chore", "🔙 revert"}
	}

	commitTypeOptions = options
}

// PromptForCommitDetails prompts the user for commit details.
func PromptForCommitDetails() (commitType string, commitScope string, commitMessage string, err error) {

	var cType string
	var cScope string
	var cMsg string

	form := huh.NewForm(
		// Type
		huh.NewGroup(
			huh.NewSelect[string]().Title("Select the commit type").Value(&cType).Options(huh.NewOptions(commitTypeOptions...)...),
		),
		// Scope
		huh.NewGroup(
			huh.NewInput().
				Value(&cScope).
				Title("Enter the commit scope (optional)"),
		),
		// Message
		huh.NewGroup(
			huh.NewInput().
				Value(&cMsg).
				Title("Enter the commit message").
				Validate(func(v string) error {
					if len(v) < 1 {
						return fmt.Errorf("commit message cannot be empty")
					}
					return nil
				}),
		),
	).WithTheme(huh.ThemeCatppuccin())

	err = form.Run()
	if err != nil {
		return "", "", "", err
	}

	return cType, cScope, cMsg, nil
}
