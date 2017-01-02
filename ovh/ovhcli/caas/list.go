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

		containersservices, err := client.ContainersServicesList()
		common.Check(err)

		if withDetails {
			contComplete := []ovh.ContainersService{}
			for _, cont := range containersservices {
				c, err := client.ContainersServiceInfo(cont.Name)
				common.Check(err)
				contComplete = append(contComplete, *c)
			}
			containersservices = contComplete
		}

		common.FormatOutputDef(containersservices)
	},
}
