package repo

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gitflow",
	Short: "Herramienta CLI para gestión de repositorios GitHub",
}

func init() {
	// Comando principal de repos
	repoCmd := &cobra.Command{
		Use:   "repo",
		Short: "Operaciones con repositorios",
	}

	repoCmd.AddCommand(CreateRepoCmd)
	repoCmd.AddCommand(DeleteCmd)
	repoCmd.AddCommand(CloneRepoCmd)
	repoCmd.AddCommand(ListCmd) // 👈 Aquí agregas el comando list

	RootCmd.AddCommand(repoCmd)
}
