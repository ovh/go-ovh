package agent

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var serviceName string
var withDetails bool
var billingAccount string

func init() {
	cmdEasyHuntingAgentList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display telephony details")

	cmdEasyHuntingAgentList.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
	cmdEasyHuntingAgentList.PersistentFlags().StringVarP(&serviceName, "serviceName", "", "", "Service Name")
}

// cmdEasyHuntingAgentList ...
var cmdEasyHuntingAgentList = &cobra.Command{
	Use:   "list",
	Short: "Hunting commands: ovhcli telephony easyhunting hunting agent list --help",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		services, err := client.TelephonyOvhPabxHuntingAgentList(billingAccount, serviceName, withDetails)
		common.Check(err)

		common.FormatOutputDef(services)
	},
}
