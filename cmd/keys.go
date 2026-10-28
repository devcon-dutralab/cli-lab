package cmd

import (
	"github.com/spf13/cobra"

	"github.com/devcon-dutralab/cli-lab/internal/keys"
)

func newKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Manage local keys",
	}
	cmd.AddCommand(newKeysAddCmd(), newKeysListCmd())
	return cmd
}

func newKeysAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add <name>",
		Short: "Create a new named key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := keys.DefaultPath()
			if err != nil {
				return err
			}
			store, err := keys.Open(path)
			if err != nil {
				return err
			}
			k, err := store.Add(args[0])
			if err != nil {
				return err
			}
			cmd.Printf("created key %q (%s)\n", k.Name, k.Address)
			return nil
		},
	}
}

func newKeysListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List local keys",
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
			list, err := store.List()
			if err != nil {
				return err
			}
			if len(list) == 0 {
				cmd.Println("no keys")
				return nil
			}
			for _, k := range list {
				cmd.Printf("%-16s %s\n", k.Name, k.Address)
			}
			return nil
		},
	}
}
