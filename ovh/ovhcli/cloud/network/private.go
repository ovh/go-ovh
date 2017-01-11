package network

import "github.com/spf13/cobra"

func init() {
	cmdCloudNetworkPrivate.AddCommand(cmdCloudNetworkPrivateShow)

}

// cmdCloudNetworkPrivate ...
var cmdCloudNetworkPrivate = &cobra.Command{
	Use:   "private",
	Short: "Network commands: ovhcli cloud network private --help",
	Long:  `Network commands: ovhcli cloud network private <command>`,
}
