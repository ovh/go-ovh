package region

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
	Use:   "region",
	Short: "Queue region commands: ovhcli dbaas queue region --help",
	Long:  `Queue region commands: ovhcli dbaas queue region <command>`,
}
