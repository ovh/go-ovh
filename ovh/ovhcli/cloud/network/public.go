package network

import "github.com/spf13/cobra"

func init() {
	cmdCloudNetworkPublic.AddCommand(cmdCloudNetworkPublicShow)

}

// cmdCloudNetworkPublic ...
var cmdCloudNetworkPublic = &cobra.Command{
	Use:   "public",
	Short: "Network commands: ovhcli cloud network public --help",
	Long:  `Network commands: ovhcli cloud network public <command>`,
}
