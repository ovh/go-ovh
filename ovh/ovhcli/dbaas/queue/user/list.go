package user

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
	Short: "List all users on a service: ovhcli dbaas queue user (--name=AppName | <--id=appID>)",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		if name == "" {
			common.WrongUsage(cmd)
		}

		app, errInfo := client.DBaasQueueAppInfoByName(name)
		common.Check(errInfo)
		id = app.ID

		apps, err := client.DBaasQueueUserList(id, withDetails)
		common.Check(err)
		common.FormatOutputDef(apps)
	},
}
