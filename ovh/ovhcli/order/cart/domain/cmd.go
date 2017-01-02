package domain

import "github.com/spf13/cobra"

var cartID string

func init() {
	Cmd.AddCommand(cmdListProductsDomain)
	Cmd.AddCommand(cmdAddProductDomain)

	Cmd.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")

}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "domain",
	Short: "Domain commands: ovhcli order cart domain --help",
	Long:  `Domain commands: ovhcli order cart domain <command>`,
}
