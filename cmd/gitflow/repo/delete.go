package repo

import (
	"fmt"
	"os"

	"github.com/EduardoNovato/gitflow/internal/app/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <nombre>",
	Short: "Elimina un repositorio de GitHub",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsAuthenticated() {
			fmt.Println("No autenticado. Ejecuta primero: gh auth login")
			os.Exit(1)
		}

		repoName := args[0]
		if !utils.ConfirmDelete(repoName) {
			fmt.Println("Operación cancelada")
			return
		}

		if err := utils.RepoDelete(repoName); err != nil {
			fmt.Printf("\nError: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nRepositorio '%s' eliminado exitosamente\n", repoName)
	},
}
