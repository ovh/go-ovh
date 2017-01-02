package hunting

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/runabove/go-sdk/ovh/ovhcli/telephony/easyhunting/hunting/agent"

	"github.com/spf13/cobra"
)

var serviceName string
var withDetails bool
var billingAccount string

func init() {
	Cmd.AddCommand(agent.Cmd)

	Cmd.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display telephony details")
	Cmd.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
	Cmd.PersistentFlags().StringVarP(&serviceName, "serviceName", "", "", "Service Name")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "hunting",
	Short: "Hunting commands: ovhcli telephony easyhunting hunting [--help] [--billingAccount=<billingAccount>]  [--serviceName=<serviceName>]",
	Long:  `Hunting commands: ovhcli telephony easyhunting hunting <command>`,
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		services, err := client.TelephonyOvhPabxHunting(billingAccount, serviceName)
		common.Check(err)

		common.FormatOutputDef(services)
	},
}
