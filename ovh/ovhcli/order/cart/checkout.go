package cart

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdCartCheckoutGet = &cobra.Command{
	Use:   "GetCheckout <cartID>",
	Short: "Get checkout cart : ovhcli order cart getCheckout <cartID>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		d, err := client.OrderGetCheckoutCart(cartID)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}

var cmdCartCheckoutPost = &cobra.Command{
	Use:   "postCheckout <cartID>",
	Short: "Post checkout cart : ovhcli order cart postCheckout <cartID>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		d, err := client.OrderPostCheckoutCart(cartID, true)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}
