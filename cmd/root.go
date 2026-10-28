// Package cmd holds cli-lab's cobra command tree.
package cmd

import "github.com/spf13/cobra"

const version = "0.1.0"

// Execute runs the root command.
func Execute() error {
	return newRootCmd().Execute()
}

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "cli-lab",
		Short:         "Community CLI for the bor-lab simulation chain",
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(newKeysCmd(), newTxCmd(), newStatusCmd())
	return root
}
