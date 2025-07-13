package repo

import (
	"fmt"
	"os"

	"github.com/EduardoNovato/gitflow/internal/app/utils"
	"github.com/spf13/cobra"
)

var (
	limit      int
	visibility string
	isFork     bool
	isSource   bool
	language   string
	owner      string
)

var ListCmd = &cobra.Command{
	Use:   "list [owner]",
	Short: "Lista repositorios del usuario o una organización",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsAuthenticated() {
			fmt.Println("No autenticado. Ejecuta primero: gh auth login")
			os.Exit(1)
		}

		if len(args) > 0 {
			owner = args[0]
		} else {
			owner = utils.GetCurrentUser()
		}

		err := utils.RepoList(owner, limit, visibility, isFork, isSource, language)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}
