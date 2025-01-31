package koth

import (
	"github.com/scorify/scorify/pkg/config"
	"github.com/sirupsen/logrus"
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

var kothCheckNames []string

func init() {
	Cmd.Flags().StringArrayVar(&kothCheckNames, "check", []string{}, "Name of koth checks")
}

func run(cmd *cobra.Command, args []string) {
	if len(kothCheckNames) == 0 {
		err := cmd.Help()
		if err != nil {
			logrus.WithError(err).Fatal("failed to print help")
		}

		return
	}

	// Retrive check configurations

	// Start koth worker

	// Handle restarts
}
