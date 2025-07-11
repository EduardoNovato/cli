package prompts

// import (
// 	"fmt"

// 	"github.com/AlecAivazis/survey/v2"
// )

// // RepoDetails recolecta información para crear un repositorio
// func RepoDetails() (name, visibility string, initReadme bool, err error) {
// 	qs := []*survey.Question{
// 		{
// 			Name: "name",
// 			Prompt: &survey.Input{
// 				Message: "Nombre del repositorio:",
// 			},
// 			Validate: survey.Required,
// 		},
// 		{
// 			Name: "visibility",
// 			Prompt: &survey.Select{
// 				Message: "Visibilidad:",
// 				Options: []string{"public", "private"},
// 				Default: "public",
// 			},
// 		},
// 		{
// 			Name: "initReadme",
// 			Prompt: &survey.Confirm{
// 				Message: "¿Inicializar con README?",
// 				Default: true,
// 			},
// 		},
// 	}

// 	answers := struct {
// 		Name       string
// 		Visibility string
// 		InitReadme bool
// 	}{}

// 	err = survey.Ask(qs, &answers)
// 	return answers.Name, answers.Visibility, answers.InitReadme, err
// }

// // ConfirmDelete solicita confirmación para eliminar un repo
// func ConfirmDelete(repoName string) bool {
// 	var confirm bool
// 	prompt := &survey.Confirm{
// 		Message: fmt.Sprintf("¿Estás seguro de eliminar '%s'?", repoName),
// 		Default: false,
// 	}
// 	survey.AskOne(prompt, &confirm)
// 	return confirm
// }
