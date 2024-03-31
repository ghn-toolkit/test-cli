// Copyright 2023 Kent Duong and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package registry

import (
	"embed"

	"github.com/AlecAivazis/survey/v2"
)

// CLIVersion version of Create Go App CLI.
const CLIVersion string = "0.1.0"

// Variables struct for Ansible variables (inventory, hosts).
type Variables struct {
	List map[string]interface{}
}

// CreateAnswers struct for a survey's answers for `create` command.
type CreateAnswers struct {
	Backend       string
	AgreeCreation bool `survey:"agree"`
}

var (
	// EmbedMiscFiles misc files and configs.
	//go:embed misc/*
	EmbedMiscFiles embed.FS

	// EmbedTemplates template files.
	//go:embed templates/*
	EmbedTemplates embed.FS

	// CreateQuestions survey's questions for `create` command.
	CreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Select{
				Message: "Choose a backend framework:",
				Options: []string{
					"fe-service",
				},
				Default:  "fe-service",
				PageSize: 1,
			},
			Validate: survey.Required,
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}

	// CustomCreateQuestions survey's questions for `create -c` command.
	CustomCreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Input{
				Message: "Enter URL to the custom backend repository:",
			},
			Validate: survey.Required,
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}
)
