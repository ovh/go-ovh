package easyhunting

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool
var billingAccount string

func init() {
	cmdEasyHuntingList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display telephony details")
	cmdEasyHuntingList.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
}

var cmdEasyHuntingList = &cobra.Command{
	Use:   "list",
	Short: "List all telephony billing account: ovhcli telephony easyhunting list --billingAccount=<billingAccount>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		services, err := client.TelephonyEasyHuntingList(billingAccount, withDetails)
		common.Check(err)

		common.FormatOutputDef(services)
	},
}
