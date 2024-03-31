// Copyright 2023 Kent Duong and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/ghn-toolkit/test-cli/pkg/cga"
	"github.com/ghn-toolkit/test-cli/pkg/registry"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(
		&useCustomTemplate,
		"template", "t", false,
		"enables to use custom backend and frontend templates",
	)
}

// createCmd represents the `create` command.
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create a new project via interactive UI",
	Long:    "\nCreate a new project via interactive UI.",
	RunE:    runCreateCmd,
}

// runCreateCmd represents runner for the `create` command.
func runCreateCmd(cmd *cobra.Command, args []string) error {
	// Start message.
	cga.ShowMessage(
		"",
		fmt.Sprintf(
			"Create a new project via Create Go App CLI v%v...",
			registry.CLIVersion,
		),
		true, true,
	)

	// Start survey.
	if useCustomTemplate {
		// Custom survey.
		if err := survey.Ask(
			registry.CustomCreateQuestions,
			&customCreateAnswers,
			survey.WithIcons(surveyIconsConfig),
		); err != nil {
			return cga.ShowError(err.Error())
		}

		// Define variables for better display.
		backend = customCreateAnswers.Backend
	} else {
		// Default survey.
		if err := survey.Ask(
			registry.CreateQuestions,
			&createAnswers,
			survey.WithIcons(surveyIconsConfig),
		); err != nil {
			return cga.ShowError(err.Error())
		}

		// Define variables for better display.
		backend = fmt.Sprintf(
			"github.com/ghn-toolkit/%v-go-template",
			strings.ReplaceAll(createAnswers.Backend, "/", "_"),
		)
	}

	// Catch the cancel action (hit "n" in the last question).
	if (!createAnswers.AgreeCreation && !useCustomTemplate) || (!customCreateAnswers.AgreeCreation && useCustomTemplate) {
		cga.ShowMessage(
			"",
			"Oh no! You said \"no\", so I won't create anything. Hope to see you soon!",
			true, true,
		)
		return nil
	}

	// Start timer.
	startTimer := time.Now()

	/*
		The project's backend part creation.
	*/

	// Clone backend files from git repository.
	if err := cga.GitClone("backend", backend); err != nil {
		return cga.ShowError(err.Error())
	}

	// Show success report.
	cga.ShowMessage(
		"success",
		fmt.Sprintf("Backend was created with template `%v`!", backend),
		true, false,
	)

	/*
		Cleanup project.
	*/

	// Stop timer.
	stopTimer := cga.CalculateDurationTime(startTimer)
	cga.ShowMessage(
		"info",
		fmt.Sprintf("Completed in %v seconds!", stopTimer),
		true, true,
	)

	cga.ShowMessage(
		"",
		"* A helpful documentation and next steps with your project is here https://github.com/ghn-toolkit/cli/wiki",
		false, true,
	)
	cga.ShowMessage(
		"",
		"Have a happy new project! :)",
		false, true,
	)

	return nil
}
