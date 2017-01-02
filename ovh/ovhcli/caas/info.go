package caas

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var containerservid string

func init() {
	cmdContainersServiceInfo.PersistentFlags().StringVarP(&containerservid, "serviceName", "", "", "Containers Service Name")
}

var cmdContainersServiceInfo = &cobra.Command{
	Use:   "info",
	Short: "Info about a project: ovhcli caas info",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		containersserviceinfo, err := client.ContainersServiceInfo(containerservid)
		common.Check(err)
		common.FormatOutputDef(containersserviceinfo)
	},
}
