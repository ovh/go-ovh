package role

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
	Use:   "role",
	Short: "Queue role commands: ovhcli dbaas queue role --help",
	Long:  `Queue role commands: ovhcli dbaas queue role <command>`,
}
