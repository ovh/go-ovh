package vrack

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdVrackList)
}

// Cmd vrack
var Cmd = &cobra.Command{
	Use:   "vrack",
	Short: "Vrack commands: ovhcli vrack --help",
	Long:  `Vrack commands: ovhcli vrack <command>`,
}
