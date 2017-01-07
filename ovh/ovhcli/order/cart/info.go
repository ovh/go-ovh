package cart

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
	cmdCartInfo.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")
}

var cmdCartInfo = &cobra.Command{
	Use:   "info <cartID>",
	Short: "Retrieve cart info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		d, err := client.OrderCartInfo(cartID)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}
