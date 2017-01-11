package key

import "github.com/spf13/cobra"

var id string
var name string

func init() {
	Cmd.AddCommand(cmdInfo)
	Cmd.AddCommand(cmdList)

	Cmd.PersistentFlags().StringVarP(&id, "id", "", "", "Your Application ID")
	Cmd.PersistentFlags().StringVarP(&name, "name", "", "", "Your Application Name")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "key",
	Short: "Queue Key commands: ovhcli dbaas queue key --help",
	Long:  `Queue Key commands: ovhcli dbaas queue key <command>`,
}
