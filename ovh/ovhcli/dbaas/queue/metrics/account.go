package metrics

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
}

var cmdAccount = &cobra.Command{
	Use:   "account",
	Short: "Get metrics account: ovhcli dbaas queue metrics account (--name=AppName | <--id=appID>)",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		apps, err := client.DBaasQueueMetricsAccount(id)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
