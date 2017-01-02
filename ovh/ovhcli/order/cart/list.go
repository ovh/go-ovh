package cart

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdCartList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display order cart details")
}

var cmdCartList = &cobra.Command{
	Use:   "list",
	Short: "List all carts: ovhcli cart list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		carts, err := client.OrderCartList()
		common.Check(err)

		if withDetails {
			carts = getDetailledCartsList(client, carts)
		}

		common.FormatOutputDef(carts)
	},
}

func getDetailledCartsList(client *ovh.Client, carts []ovh.OrderCart) []ovh.OrderCart {

	cartsChan, errChan := make(chan ovh.OrderCart), make(chan error)
	for _, cart := range carts {
		go func(cart ovh.OrderCart) {
			c, err := client.OrderCartInfo(cart.CartID)
			if err != nil {
				errChan <- err
				return
			}
			cartsChan <- *c
		}(cart)
	}

	cartsComplete := []ovh.OrderCart{}

	for i := 0; i < len(carts); i++ {
		select {
		case cart := <-cartsChan:
			cartsComplete = append(cartsComplete, cart)
		case err := <-errChan:
			common.Check(err)
		}
	}

	return cartsComplete
}
