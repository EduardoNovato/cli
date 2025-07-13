package repo

import (
	"fmt"
	"os"

	"github.com/EduardoNovato/gitflow/internal/app/utils"
	"github.com/spf13/cobra"
)

var ArchiveRepoCmd = &cobra.Command{
	Use:   "archive [<repository>]",
	Short: "Archiva un repositorio de GitHub",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		if !utils.ConfirmRepoArchive(repoName) {
			fmt.Println("Operación cancelada")
			return
		}

		if err := utils.RepoArchive(repoName); err != nil {
			fmt.Printf("\nError: %v\n", err)
			fmt.Println("Sugerencia: Si persiste el problema, intenta:")
			fmt.Println("1. Ejecutar manualmente: gh auth login")
			fmt.Println("2. Verificar tu conexión a internet")
			os.Exit(1)
		}
	},
}
