package topic

import "github.com/spf13/cobra"

var id string
var name string

func init() {
	Cmd.AddCommand(cmdList)

	Cmd.PersistentFlags().StringVarP(&id, "id", "", "", "Your Application ID")
	Cmd.PersistentFlags().StringVarP(&name, "name", "", "", "Your Application Name")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "topic",
	Short: "Queue topic commands: ovhcli dbaas queue topic --help",
	Long:  `Queue topic commands: ovhcli dbaas queue topic <command>`,
}
