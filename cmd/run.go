package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start a webserver",
	// TODO: Add a description here
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Provide subcommand")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().IntP("port", "p", 9000, "Port number to listen on")
}
