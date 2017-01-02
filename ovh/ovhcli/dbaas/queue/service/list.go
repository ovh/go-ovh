package service

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display details")
}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List all services: ovhcli dbaas queue service list (--name=AppName | <--id=appID>)",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		apps, err := client.DBaasQueueAppList(withDetails)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
