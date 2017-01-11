package network

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdCloudNetworkPublic)
	Cmd.AddCommand(cmdCloudNetworkPrivate)

}

// cmdCloudNetwork ...
var Cmd = &cobra.Command{
	Use:     "network",
	Short:   "Network commands: ovhcli cloud network --help",
	Long:    `Network commands: ovhcli cloud network <command>`,
	Aliases: []string{"net"},
}
