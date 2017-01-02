package domain

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/spf13/cobra"
)

var duration string
var offerID string
var quantity int

func init() {

	cmdAddProductDomain.PersistentFlags().StringVarP(&duration, "duration", "d", "P1Y", "domain")
	cmdAddProductDomain.PersistentFlags().StringVarP(&offerID, "offerID", "o", "", "domain")
	cmdAddProductDomain.PersistentFlags().IntVarP(&quantity, "quantity", "q", 1, "domain")

}

var cmdListProductsDomain = &cobra.Command{
	Use:   "list <domain>",
	Short: "Get list products about a domain name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		domain := args[0]

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		c, err := client.OrderGetProductsDomain(cartID, domain)
		common.Check(err)
		common.FormatOutputDef(c)
	},
}

var cmdAddProductDomain = &cobra.Command{
	Use:   "add <domain>",
	Short: "Add domain product into the cart",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		domain := args[0]

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		c, err := client.OrderAddProductDomain(cartID, ovh.OrderPostDomainReq{
			Domain:   domain,
			Duration: duration,
			OfferID:  offerID,
			Quantity: quantity,
		})
		common.Check(err)
		common.FormatOutputDef(c)
	},
}
