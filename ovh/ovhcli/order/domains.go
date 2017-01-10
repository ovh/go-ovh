package order

import (
	"fmt"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/runabove/go-sdk/ovh/types"
	"github.com/spf13/cobra"
)

var withOffer string
var withConfigs string

var withOwner string

func init() {
	CmdDomain.PersistentFlags().StringVarP(&withOffer, "withOffer", "", "gold", "offer on your domain (gold, diamond, platinium)")
	CmdDomain.PersistentFlags().StringVarP(&withConfigs, "withConfigs", "", "", "configs file")

	CmdDomainTrade.PersistentFlags().StringVarP(&withOwner, "withOwner", "", "", "configs file")

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

		cart, err := client.OrderCreateCart(types.OrderCartPost{OvhSubsidiary: "FR"})
		common.Check(err)

		err = client.OrderAssignCart(cart.CartID)
		common.Check(err)

		// TODO products, err := client.OrderGetProductsDomain(cart.CartID, domain)
		_, err = client.OrderGetProductsDomain(cart.CartID, domain)
		common.Check(err)

		// var chooseProduct string
		common.Check(fmt.Errorf("product.Offer -> not in OrderCartProductInformation"))
		/*for _, product := range products {
			if product.Offer == withOffer {
				chooseProduct = product.OfferID
				break
			}
		}
		if chooseProduct == "" {
			err = fmt.Errorf("Cannot find product for domain %s and this offer %s", domain, withOffer)
			common.Check(err)
		}
		_, err = client.OrderAddProductDomain(cart.CartID, types.OrderCartDomainPost{
			Domain:   domain,
			Duration: "P1Y",
			OfferID:  chooseProduct,
			Quantity: 1,
		})
		common.Check(err)

		order, err := client.OrderPostCheckoutCart(cart.CartID, false)
		common.Check(err)
		common.FormatOutputDef(order)
		*/

	},
}

// CmdDomainTrade in order domain trade
var CmdDomainTrade = &cobra.Command{
	Use:   "domainTrade <domain>",
	Short: "Order domain trade",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		domain := args[0]

		if withOwner == "" {
			common.Exit("Missing withOwner option")
		}

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		cart, err := client.OrderCreateCart(types.OrderCartPost{OvhSubsidiary: "FR"})
		common.Check(err)

		err = client.OrderAssignCart(cart.CartID)
		common.Check(err)

		serviceOptions, err := client.OrderGetCartServiceOptions(domain)
		common.Check(err)

		var serviceOptionChoosed *types.OrderCartGenericOptionDefinition
		for i := range serviceOptions {
			serviceOption := serviceOptions[i]
			if serviceOption.Family == "trade" {
				serviceOptionChoosed = &serviceOption
				break
			}
		}
		if serviceOptionChoosed == nil {
			common.Exit("Cannot find service options of type trade for domain %s", domain)
		}

		item, err := client.OrderAddCartServiceOption(domain, types.OrderCartServiceOptionDomainPost{
			Duration:    "P1Y", //always P1Y for a trade
			Quantity:    1,     //always 1 for quantity for a trade
			CartID:      cart.CartID,
			PlanCode:    serviceOptionChoosed.PlanCode,
			PricingMode: "default",
		})
		common.Check(err)

		_, err = client.OrderCartAddConfiguration(cart.CartID, item.ItemID, "OWNER_CONTACT", withOwner)
		common.Check(err)

		order, err := client.OrderPostCheckoutCart(cart.CartID, true)
		common.Check(err)

		common.FormatOutputDef(order)

	},
}
