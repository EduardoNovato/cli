package gitflow

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitflow",
	Short: "Herramienta CLI para simplificar flujos de trabajo con GitHub",
	Long:  `GitFlow es una herramienta que simplifica la interacción con GitHub mediante comandos interactivos.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
