package cart

import (
	"fmt"
	"os"
	"strconv"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var configWithDetails bool

var label string
var value string

func init() {
	CmdCartItemConfigurationsList.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")
	CmdCartItemConfigurationsList.PersistentFlags().BoolVarP(&configWithDetails, "withDetails", "", false, "Display domain details")

	CmdCartItemConfigurationInfo.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")

	CmdCartItemConfigurationAdd.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")
	CmdCartItemConfigurationAdd.PersistentFlags().StringVarP(&label, "withLabel", "", "", "Label of config")
	CmdCartItemConfigurationAdd.PersistentFlags().StringVarP(&value, "withValue", "", "", "Value of config")

	CmdCartItemConfigurationRemove.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")

	CmdCartItemRequiredConfigurations.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")

}

//CmdCartItemConfigurationsList list all configurations for an item
var CmdCartItemConfigurationsList = &cobra.Command{
	Use:   "listConfigs",
	Short: "List all configs of an item: ovhcli order cart listConfigs <itemId>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		client, err := ovh.NewDefaultClient()

		common.Check(err)
		itemID := args[0]
		i, err := strconv.Atoi(itemID)
		common.Check(err)

		configs, err := client.OrderCartConfigurationsList(cartID, i)
		common.Check(err)

		if configWithDetails {
			configs = getDetailledConfigurationsList(client, i, configs)
		}

		common.FormatOutputDef(configs)
	},
}

//CmdCartItemConfigurationInfo get configuration info for an item
var CmdCartItemConfigurationInfo = &cobra.Command{
	Use:   "config <itemId> <configId>",
	Short: "Get config's info of a cart: ovhcli order cart config <itemId> <configId>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			common.WrongUsage(cmd)
		}
		itemID := args[0]
		i, err := strconv.Atoi(itemID)
		common.Check(err)

		configID := args[1]
		c, err := strconv.Atoi(configID)
		common.Check(err)

		client, err := ovh.NewDefaultClient()
		common.Check(err)
		config, err := client.OrderCartConfigurationInfo(cartID, i, c)
		common.Check(err)

		common.FormatOutputDef(config)
	},
}

//CmdCartItemConfigurationAdd add configuration on an item
var CmdCartItemConfigurationAdd = &cobra.Command{
	Use:   "addConfig <itemId>",
	Short: "Add config on an item: ovhcli order cart addConfig <itemId>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		itemID := args[0]
		i, err := strconv.Atoi(itemID)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		item, err := client.OrderCartAddConfiguration(cartID, i, label, value)
		common.Check(err)

		common.FormatOutputDef(item)
	},
}

//CmdCartItemConfigurationRemove remove config on an item
var CmdCartItemConfigurationRemove = &cobra.Command{
	Use:   "removeConfig <itemId> <configId>",
	Short: "Delete item's info of a cart: ovhcli order cart removeConfig <itemId> <configId>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			common.WrongUsage(cmd)
		}
		itemID := args[0]
		i, err := strconv.Atoi(itemID)
		common.Check(err)

		configID := args[1]
		c, err := strconv.Atoi(configID)
		common.Check(err)

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		config, err := client.OrderCartDeleteConfiguration(cartID, i, c)
		common.Check(err)

		common.FormatOutputDef(config)
	},
}

func getDetailledConfigurationsList(client *ovh.Client, itemID int, configs []ovh.OrderCartConfigurationItem) []ovh.OrderCartConfigurationItem {

	resChan, errChan := make(chan ovh.OrderCartConfigurationItem), make(chan error)
	for _, config := range configs {
		go func(config ovh.OrderCartConfigurationItem) {
			i, err := client.OrderCartConfigurationInfo(cartID, itemID, config.ID)
			if err != nil {
				errChan <- err
				return
			}
			resChan <- *i
		}(config)
	}

	itemsComplete := []ovh.OrderCartConfigurationItem{}

	for i := 0; i < len(configs); i++ {
		select {
		case item := <-resChan:
			itemsComplete = append(itemsComplete, item)
		case err := <-errChan:
			common.Check(err)
		}
	}

	return itemsComplete
}

//CmdCartItemRequiredConfigurations list all configurations for an item
var CmdCartItemRequiredConfigurations = &cobra.Command{
	Use:   "listRequiredConfigs <itemId>",
	Short: "List all required configurations for an item: ovhcli order cart listRequiredConfigs <itemId>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		client, err := ovh.NewDefaultClient()

		common.Check(err)
		itemID := args[0]
		i, err := strconv.Atoi(itemID)
		common.Check(err)

		configs, err := client.OrderCartRequiredConfigurations(cartID, i)
		common.Check(err)

		common.FormatOutputDef(configs)
	},
}
