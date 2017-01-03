package caas

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdContainersServicesList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display containers details")
}

var cmdContainersServicesList = &cobra.Command{
	Use:   "list",
	Short: "List all containers services: ovhcli caas list",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		containersservices, err := client.ContainersServicesList(withDetails)
		common.Check(err)

		common.FormatOutputDef(containersservices)
	},
}
