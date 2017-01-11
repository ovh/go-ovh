package cart

import (
	"github.com/runabove/go-sdk/ovh/ovhcli/order/cart/domain"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdCartList)
	Cmd.AddCommand(cmdCartInfo)
	Cmd.AddCommand(cmdCartAssign)
	Cmd.AddCommand(cmdCartCreate)
	Cmd.AddCommand(cmdCartDelete)

	Cmd.AddCommand(cmdCartSummary)
	Cmd.AddCommand(cmdCartCheckoutGet)
	Cmd.AddCommand(cmdCartCheckoutPost)

	Cmd.AddCommand(CmdCartListItems)
	Cmd.AddCommand(CmdCartInfoItem)
	Cmd.AddCommand(CmdCartUpdateItem)
	Cmd.AddCommand(CmdCartDeleteItem)

	Cmd.AddCommand(CmdCartItemConfigurationsList)
	Cmd.AddCommand(CmdCartItemConfigurationInfo)
	Cmd.AddCommand(CmdCartItemConfigurationAdd)
	Cmd.AddCommand(CmdCartItemConfigurationRemove)
	Cmd.AddCommand(CmdCartItemRequiredConfigurations)

	Cmd.AddCommand(domain.Cmd)

}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "cart",
	Short: "Cart commands: ovhcli order cart --help",
	Long:  `Cart commands: ovhcli order cart <command>`,
}
