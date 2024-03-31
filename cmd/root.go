// Copyright 2023 Kent Duong and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/ghn-toolkit/test-cli/pkg/cga"
	"github.com/ghn-toolkit/test-cli/pkg/registry"
)

var (
	backend                            string                 // define project variables
	useCustomTemplate                  bool                   // define custom templates
	createAnswers, customCreateAnswers registry.CreateAnswers // define answers variable for `create` command

	// Config for survey icons and colors.
	// See: https://github.com/mgutz/ansi#style-format
	surveyIconsConfig = func(icons *survey.IconSet) {
		icons.Question.Format = "cyan"
		icons.Question.Text = "[?]"
		icons.Help.Format = "blue"
		icons.Help.Text = "Help ->"
		icons.Error.Format = "yellow"
		icons.Error.Text = "Note ->"
	}
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "cga",
	Version: registry.CLIVersion,
	Short:   "A powerful CLI for the Create Go App project",
	Long: `
A powerful CLI for the Create Go App project.

Create a new production-ready project with backend (Golang) by running CLI command.

-> Focus on writing code and thinking of business logic!
<- The Create Go App CLI will take care of the rest.

A helpful documentation and next steps -> https://github.com/ghn-toolkit/cli/wiki`,
}

func init() {
	rootCmd.SetOut(cga.Stdout)
	rootCmd.SetErr(cga.Stderr)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = rootCmd.Execute()
}
