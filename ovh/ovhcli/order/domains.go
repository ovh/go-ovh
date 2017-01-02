package order

import (
	"fmt"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/spf13/cobra"
)

var withOffer string
var withConfigs string

func init() {

	CmdDomain.PersistentFlags().StringVarP(&withOffer, "withOffer", "", "gold", "offer on your domain (gold, diamond, platinium)")
	CmdDomain.PersistentFlags().StringVarP(&withConfigs, "withConfigs", "", "", "configs file")

}

// CmdDomain order domain
var CmdDomain = &cobra.Command{
	Use:   "domain <domain>",
	Short: "Order domain",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		domain := args[0]

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		cart, err := client.OrderCreateCart(ovh.OrderCartCreateReq{OVHSubsidiary: "FR"})
		common.Check(err)

		err = client.OrderAssignCart(cart.CartID)
		common.Check(err)

		products, err := client.OrderGetProductsDomain(cart.CartID, domain)
		common.Check(err)

		var chooseProduct string
		for _, product := range products {
			if product.Offer == withOffer {
				chooseProduct = product.OfferID
				break
			}
		}
		if chooseProduct == "" {
			err = fmt.Errorf("Cannot find product for domain %s and this offer %s", domain, withOffer)
			common.Check(err)
		}
		_, err = client.OrderAddProductDomain(cart.CartID, ovh.OrderPostDomainReq{
			Domain:   domain,
			Duration: "P1Y",
			OfferID:  chooseProduct,
			Quantity: 1,
		})
		common.Check(err)

		order, err := client.OrderPostCheckoutCart(cart.CartID, false)
		common.Check(err)
		common.FormatOutputDef(order)

	},
}
