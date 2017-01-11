package order

import (
	"github.com/runabove/go-sdk/ovh/ovhcli/order/cart"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cart.Cmd)
	Cmd.AddCommand(CmdDomain)
	Cmd.AddCommand(CmdDomainTrade)

}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "order",
	Short: "Order commands: ovhcli order --help",
	Long:  `Order commands: ovhcli order <command>`,
}
