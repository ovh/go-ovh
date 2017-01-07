package cart

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/runabove/go-sdk/ovh/types"

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

		client, errc := ovh.NewDefaultClient()
		common.Check(errc)

		carts, errl := client.OrderCartList()
		common.Check(errl)

		if withDetails {
			carts = getDetailledCartsList(client, carts)
		}

		common.FormatOutputDef(carts)
	},
}

func getDetailledCartsList(client *ovh.Client, carts []types.OrderCart) []types.OrderCart {

	cartsChan, errChan := make(chan types.OrderCart), make(chan error)
	for _, cart := range carts {
		go func(cart types.OrderCart) {
			c, err := client.OrderCartInfo(cart.CartID)
			if err != nil {
				errChan <- err
				return
			}
			cartsChan <- *c
		}(cart)
	}

	cartsComplete := []types.OrderCart{}

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
