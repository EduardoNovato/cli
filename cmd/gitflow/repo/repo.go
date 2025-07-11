package repo

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gitflow",
	Short: "Herramienta CLI para gestión de repositorios GitHub",
}

func init() {
	// Configuración de comandos
	repoCmd := &cobra.Command{
		Use:   "repo",
		Short: "Operaciones con repositorios",
	}

	repoCmd.AddCommand(CreateRepoCmd)
	repoCmd.AddCommand(DeleteCmd)
	RootCmd.AddCommand(repoCmd)
}
