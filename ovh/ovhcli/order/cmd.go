package order

import (
	"github.com/runabove/go-sdk/ovh/ovhcli/order/cart"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cart.Cmd)
	Cmd.AddCommand(CmdDomain)
}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "order",
	Short: "order commands: ovhcli order --help",
	Long:  `order commands: ovhcli order <command>`,
}
