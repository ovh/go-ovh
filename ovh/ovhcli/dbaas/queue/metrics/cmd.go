package metrics

import "github.com/spf13/cobra"

var id string
var name string

func init() {
	Cmd.AddCommand(cmdAccount)

	Cmd.PersistentFlags().StringVarP(&id, "id", "", "", "Your Application ID")
	Cmd.PersistentFlags().StringVarP(&name, "name", "", "", "Your Application Name")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "metrics",
	Short: "Queue Key commands: ovhcli dbaas queue metrics --help",
	Long:  `Queue Key commands: ovhcli dbaas queue metrics <command>`,
}
