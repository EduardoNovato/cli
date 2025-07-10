package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

func AskRepoDetails() (name string, visibility string, initReadme bool, err error) {
	qs := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Nombre del repositorio:",
			},
			Validate: survey.Required,
		},
		{
			Name: "visibility",
			Prompt: &survey.Select{
				Message: "Visibilidad:",
				Options: []string{"public", "private"},
				Default: "public",
			},
		},
		{
			Name: "initReadme",
			Prompt: &survey.Confirm{
				Message: "¿Inicializar con README?",
				Default: true,
			},
		},
	}

	answers := struct {
		Name       string
		Visibility string
		InitReadme bool
	}{}

	err = survey.Ask(qs, &answers)
	return answers.Name, answers.Visibility, answers.InitReadme, err
}
