package cmd

import (
	"github.com/spf13/cobra"

	"github.com/devcon-dutralab/cli-lab/internal/coin"
	"github.com/devcon-dutralab/cli-lab/internal/keys"
)

func newTxCmd() *cobra.Command {
	tx := &cobra.Command{Use: "tx", Short: "Transaction subcommands"}
	tx.AddCommand(newTxSendCmd())
	return tx
}

func newTxSendCmd() *cobra.Command {
	var from, to, amount string

	send := &cobra.Command{
		Use:   "send",
		Short: "Send tokens from one account to another",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			c, err := coin.ParseCoin(amount)
			if err != nil {
				return err
			}

			path, err := keys.DefaultPath()
			if err != nil {
				return err
			}
			store, err := keys.Open(path)
			if err != nil {
				return err
			}
			if _, err := store.Get(from); err != nil {
				return err
			}

			cmd.Printf("broadcast send: %s -> %s amount %s\n", from, to, c)
			return nil
		},
	}
	send.Flags().StringVar(&from, "from", "", "sender key name")
	send.Flags().StringVar(&to, "to", "", "recipient address")
	send.Flags().StringVar(&amount, "amount", "", "amount, e.g. 100ubor")
	return send
}
