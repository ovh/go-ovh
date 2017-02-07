package user

import (
	"github.com/spf13/cobra"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
)

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Get User Info: ovhcli dbaas queue user info (--name=AppName | <--id=appID>) --user=username",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		if name == "" {
			common.WrongUsage(cmd)
		}

		if userName == "" {
			common.WrongUsage(cmd)
		}

		app, errInfo := client.DBaasQueueAppInfoByName(name)
		common.Check(errInfo)
		id = app.ID

		users, errUsers := client.DBaasQueueUserList(id, true)
		common.Check(errUsers)
		for _, u := range users {
			if u.Name == userName {
				userID = u.ID
			}
		}

		checkUser()

		user, err := client.DBaasQueueUserInfo(id, userID)
		common.Check(err)

		common.FormatOutputDef(user)
	},
}
