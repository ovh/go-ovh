package agent

import (
	"github.com/spf13/cobra"
)

func init() {

	Cmd.AddCommand(cmdEasyHuntingAgentUpdate)
	Cmd.AddCommand(cmdEasyHuntingAgentList)
	Cmd.AddCommand(cmdEasyHuntingAgentInfo)
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "agent",
	Short: "ovhcli telephony easyhunting hunting agent --help",
	Long:  `ovhcli telephony easyhunting hunting agent <command>`,
}
