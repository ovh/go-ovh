package agent

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var (
	wrapUpTime        int
	number            string
	timeout           int
	status            string
	simultaneousLines int
	breakStatus       int
)

func init() {
	cmdEasyHuntingAgentUpdate.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
	cmdEasyHuntingAgentUpdate.PersistentFlags().StringVarP(&serviceName, "serviceName", "", "", "Service Name")
	cmdEasyHuntingAgentUpdate.PersistentFlags().IntVarP(&agentID, "agentID", "", 0, "Agent ID")

	cmdEasyHuntingAgentUpdate.PersistentFlags().IntVarP(&wrapUpTime, "wrapUpTime", "", 0, "The wrap up time (in seconds) after the calls")
	cmdEasyHuntingAgentUpdate.PersistentFlags().StringVarP(&number, "number", "", "", "The number of the agent")
	cmdEasyHuntingAgentUpdate.PersistentFlags().IntVarP(&timeout, "timeout", "", 0, "The waiting timeout (in seconds) before hangup an assigned called")
	cmdEasyHuntingAgentUpdate.PersistentFlags().StringVarP(&status, "status", "", "", "The current status of the agent")
	cmdEasyHuntingAgentUpdate.PersistentFlags().IntVarP(&simultaneousLines, "simultaneousLines", "", 0, "The maximum of simultaneous calls that the agent will receive from the hunting")
	cmdEasyHuntingAgentUpdate.PersistentFlags().IntVarP(&breakStatus, "breakStatus", "", 0, "The id of the current break status of the agent")
}

var cmdEasyHuntingAgentUpdate = &cobra.Command{
	Use:   "update",
	Short: "Get info on a easyhunting: ovhcli telephony easyhunting hunting agent info --billingAccount=aa --serviceName=bb --agentID=cc",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		a := ovh.TelephonyOvhPabxHuntingAgent{
			WrapUpTime:        wrapUpTime,
			Number:            number,
			Timeout:           timeout,
			Status:            status,
			SimultaneousLines: simultaneousLines,
			BreakStatus:       breakStatus,
		}

		d, err := client.TelephonyOvhPabxHuntingAgentUpdate(billingAccount, serviceName, agentID, a)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}
