package cart

import (
	"fmt"
	"os"
	"strconv"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/runabove/go-sdk/ovh/types"

	"github.com/spf13/cobra"
)

var itemsWithDetails bool
var cartID string

var duration string
var quantity int

func init() {
	CmdCartListItems.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")
	CmdCartListItems.PersistentFlags().BoolVarP(&itemsWithDetails, "withDetails", "", false, "Display domain details")

	CmdCartInfoItem.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")

	CmdCartUpdateItem.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")
	CmdCartUpdateItem.PersistentFlags().StringVarP(&duration, "duration", "", "", "duration of your item")
	CmdCartUpdateItem.PersistentFlags().IntVarP(&quantity, "quantity", "", 1, "quantity of item")

	CmdCartDeleteItem.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")

}

//CmdCartListItems list all item of a cart
var CmdCartListItems = &cobra.Command{
	Use:   "listItems",
	Short: "List all items of a cart: ovhcli order cart list",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		items, err := client.OrderCartItemList(cartID)
		common.Check(err)

		if itemsWithDetails {
			items = getDetailledItemList(client, items)
		}

		common.FormatOutputDef(items)
	},
}

//CmdCartInfoItem get item info of a cart
var CmdCartInfoItem = &cobra.Command{
	Use:   "item <itemId>",
	Short: "Get item's info of a cart: ovhcli order cart item <itemId>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		itemID := args[0]
		i, err := strconv.ParseInt(itemID, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		client, err := ovh.NewDefaultClient()
		common.Check(err)
		item, err := client.OrderCartItemInfo(cartID, i)
		common.Check(err)

		common.FormatOutputDef(item)
	},
}

//CmdCartUpdateItem update item of a cart
var CmdCartUpdateItem = &cobra.Command{
	Use:   "updateItem <itemId>",
	Short: "Update item's info of a cart: ovhcli order cart updateItem <itemId>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		itemID := args[0]
		i, err := strconv.ParseInt(itemID, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		item, err := client.OrderUpdateCartItem(cartID, i, duration, quantity)
		common.Check(err)

		common.FormatOutputDef(item)
	},
}

//CmdCartDeleteItem delete item of a cart
var CmdCartDeleteItem = &cobra.Command{
	Use:   "deleteItem <itemId>",
	Short: "Delete item's info of a cart: ovhcli order cart deleteItem <itemId>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		itemID := args[0]
		i, err := strconv.ParseInt(itemID, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		item, err := client.OrderDeleteCartItem(cartID, i)
		common.Check(err)

		common.FormatOutputDef(item)
	},
}

func getDetailledItemList(client *ovh.Client, items []types.OrderCartItem) []types.OrderCartItem {

	itemsChan, errChan := make(chan types.OrderCartItem), make(chan error)
	for _, item := range items {
		go func(item types.OrderCartItem) {
			i, err := client.OrderCartItemInfo(cartID, item.ItemID)
			if err != nil {
				errChan <- err
				return
			}
			itemsChan <- *i
		}(item)
	}

	itemsComplete := []types.OrderCartItem{}

	for i := 0; i < len(items); i++ {
		select {
		case item := <-itemsChan:
			itemsComplete = append(itemsComplete, item)
		case err := <-errChan:
			common.Check(err)
		}
	}

	return itemsComplete
}
