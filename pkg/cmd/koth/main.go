package koth

import (
	"github.com/scorify/scorify/pkg/config"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "koth",
	Short:   "Start scoring koth worker",
	Long:    "Start scoring koth worker",
	Aliases: []string{"k", "worker", "w"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitKoth()
	},

	Run: run,
}

func run(cmd *cobra.Command, args []string) {
}
