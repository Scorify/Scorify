package cmd

import (
	"github.com/scorify/scorify/pkg/cmd/koth"
	"github.com/scorify/scorify/pkg/cmd/minion"
	"github.com/scorify/scorify/pkg/cmd/server"
	"github.com/scorify/scorify/pkg/cmd/setup"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scorify",
	Short: "Scorify is a scoring engine for Cybersecurity Competitions",
	Long:  "Scorify is a scoring engine for Cybersecurity Competitions",
	Run:   run,
}

// print help if no subcommand is given
func run(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		logrus.WithError(err).Fatal("failed to print help")
	}
}

// Entrypoint for all commands
func Execute() error {
	return rootCmd.Execute()
}

// registers all commands
func init() {
	rootCmd.AddCommand(
		minion.Cmd,
		server.Cmd,
		setup.Cmd,
		koth.Cmd,
	)
}
