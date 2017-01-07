package agent

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/runabove/go-sdk/ovh/types"

	"github.com/spf13/cobra"
)

var (
	wrapUpTime        int64
	number            string
	timeout           int64
	status            string
	simultaneousLines int64
	breakStatus       int64
)

func init() {
	cmdEasyHuntingAgentUpdate.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
	cmdEasyHuntingAgentUpdate.PersistentFlags().StringVarP(&serviceName, "serviceName", "", "", "Service Name")
	cmdEasyHuntingAgentUpdate.PersistentFlags().Int64VarP(&agentID, "agentID", "", 0, "Agent ID")

	cmdEasyHuntingAgentUpdate.PersistentFlags().Int64VarP(&wrapUpTime, "wrapUpTime", "", 0, "The wrap up time (in seconds) after the calls")
	cmdEasyHuntingAgentUpdate.PersistentFlags().StringVarP(&number, "number", "", "", "The number of the agent")
	cmdEasyHuntingAgentUpdate.PersistentFlags().Int64VarP(&timeout, "timeout", "", 0, "The waiting timeout (in seconds) before hangup an assigned called")
	cmdEasyHuntingAgentUpdate.PersistentFlags().StringVarP(&status, "status", "", "", "The current status of the agent")
	cmdEasyHuntingAgentUpdate.PersistentFlags().Int64VarP(&simultaneousLines, "simultaneousLines", "", 0, "The maximum of simultaneous calls that the agent will receive from the hunting")
	cmdEasyHuntingAgentUpdate.PersistentFlags().Int64VarP(&breakStatus, "breakStatus", "", 0, "The id of the current break status of the agent")
}

var cmdEasyHuntingAgentUpdate = &cobra.Command{
	Use:   "update",
	Short: "Get info on a easyhunting: ovhcli telephony easyhunting hunting agent info --billingAccount=aa --serviceName=bb --agentID=cc",
	Run: func(cmd *cobra.Command, args []string) {
		client, errc := ovh.NewDefaultClient()
		common.Check(errc)

		a := types.TelephonyOvhPabxHuntingAgent{
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
