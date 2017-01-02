package role

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var roleID string

func init() {
	cmdInfo.PersistentFlags().StringVarP(&roleID, "role-id", "", "", "Role ID")
}

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Get Role Info: ovhcli dbaas queue role info (--name=AppName | <--id=appID>) --role-id=roleid",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		role, err := client.DBaasQueueRoleInfo(id, roleID)
		common.Check(err)

		common.FormatOutputDef(role)
	},
}
