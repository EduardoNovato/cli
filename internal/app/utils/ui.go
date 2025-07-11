package utils

import (
	"github.com/AlecAivazis/survey/v2"
)

func Input(message string) string {
	var response string
	prompt := &survey.Input{
		Message: message,
	}
	survey.AskOne(prompt, &response)
	return response
}

func Select(message string, options []string) string {
	var response string
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}
	survey.AskOne(prompt, &response)
	return response
}

func Confirm(message string, defaultOption bool) bool {
	var response bool
	prompt := &survey.Confirm{
		Message: message,
		Default: defaultOption,
	}
	survey.AskOne(prompt, &response)
	return response
}
