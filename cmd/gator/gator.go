package main

import (
	"os"

	"github.com/open-policy-agent/gatekeeper/cmd/gator/verify"
	"github.com/spf13/cobra"
)

const version = "alpha"

func init() {
	rootCmd.AddCommand(verify.Cmd)
}

var rootCmd = &cobra.Command{
	Use:     "gator subcommand",
	Short:   "gator is a suite of authorship tools for Gatekeeper",
	Version: version,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
