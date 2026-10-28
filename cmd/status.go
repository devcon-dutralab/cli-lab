package cmd

import (
	"github.com/spf13/cobra"

	"github.com/devcon-dutralab/cli-lab/internal/keys"
)

func newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Print the CLI and local keystore status",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			path, err := keys.DefaultPath()
			if err != nil {
				return err
			}
			store, err := keys.Open(path)
			if err != nil {
				return err
			}

			list, _ := store.List()

			cmd.Printf("cli-lab %s\n", version)
			cmd.Printf("keystore: %s\n", path)
			cmd.Printf("keys: %d\n", len(list))
			return nil
		},
	}
}
