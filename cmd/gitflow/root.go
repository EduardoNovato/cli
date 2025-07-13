package gitflow

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/EduardoNovato/gitflow/internal/app/workflows"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gitflow",
	Short: "Interfaz interactiva para GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			action := selectMainAction()

			switch action {
			case "Clonaer Repositorio":
				workflows.CloneRepoFlow()
			case "Lista Repositorios":
				workflows.ListReposFlow()
			case "Crear Repositorio":
				workflows.CreateRepoFlow()
			case "Archivar Repositorio":
				workflows.ArchiveRepoFlow()
			case "Eliminar Repositorio":
				workflows.DeleteRepoFlow()
			case "Salir":
				fmt.Println("¡Hasta luego!")
				os.Exit(0)
			}
		}
	},
}

func selectMainAction() string {
	var action string
	prompt := &survey.Select{
		Message: "¿Qué deseas hacer?",
		Options: []string{
			"Clonar Repositorio",
			"Lista Repositorios",
			"Crear Repositorio",
			"Archivar Repositorio",
			"Eliminar Repositorio",
			"Salir",
		},
	}
	survey.AskOne(prompt, &action)
	return action
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
