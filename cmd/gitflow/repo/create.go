package repo

import (
	"fmt"
	"os"

	"github.com/EduardoNovato/gitflow/internal/app/utils"
	"github.com/spf13/cobra"
)

var CreateRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Crea un nuevo repositorio en GitHub",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	name, visibility, initReadme, err := prompts.AskRepoDetails()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if err := commands.CreateRepo(name, visibility, initReadme); err != nil {
	// 		panic(err)
	// 	}
	// },
	Run: func(cmd *cobra.Command, args []string) {
		name, visibility, initReadme, err := utils.AskRepoDetails()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if err := utils.RepoCreate(name, visibility, initReadme); err != nil {
			fmt.Printf("\nError: %v\n", err)
			fmt.Println("Sugerencia: Si persiste el problema, intenta:")
			fmt.Println("1. Ejecutar manualmente: gh auth login")
			fmt.Println("2. Verificar tu conexión a internet")
			os.Exit(1)
		}
	},
}
