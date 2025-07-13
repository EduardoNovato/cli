package repo

import (
	"fmt"
	"os"

	"github.com/EduardoNovato/gitflow/internal/app/utils"
	"github.com/spf13/cobra"
)

var CloneRepoCmd = &cobra.Command{
	Use:   "clone <url> [<nombre>]",
	Short: "Clona un repositorio de GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		repoURL, err := utils.AskRepoURL()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if err := utils.RepoClone(repoURL); err != nil {
			fmt.Printf("\nError: %v\n", err)
			fmt.Println("Sugerencia: Si persiste el problema, intenta:")
			fmt.Println("1. Ejecutar manualmente: gh auth login")
			fmt.Println("2. Verificar tu conexión a internet")
			os.Exit(1)
		}
	},
}
