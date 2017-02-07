package telephony

import (
	"github.com/runabove/go-sdk/ovh/ovhcli/telephony/easyhunting"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdServiceList)

	Cmd.AddCommand(easyhunting.Cmd)
}

// Cmd telephony
var Cmd = &cobra.Command{
	Use:   "telephony",
	Short: "Telephony commands: ovhcli telephony --help",
	Long:  `Telephony commands: ovhcli telephony <command>`,
}
