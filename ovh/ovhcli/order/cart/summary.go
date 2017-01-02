package cart

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdCartSummary = &cobra.Command{
	Use:   "summary <cartID>",
	Short: "Retrieve cart info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		d, err := client.OrderSummaryCart(cartID)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}
