package user

import (
	"github.com/spf13/cobra"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
)

var cmdChangePassword = &cobra.Command{
	Use:   "changepassword",
	Short: "Change password for the given user (--name=AppName) (--user=UserName)",
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
		for _, user := range users {
			if user.Name == userName {
				userID = user.ID
				break
			}
		}

		checkUser()

		user, err := client.DBaasQueueUserChangePassword(id, userID)
		common.Check(err)

		common.FormatOutputDef(user)
	},
}
