package instance

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdInstanceDelete)
	Cmd.AddCommand(cmdInstanceCreate)
	Cmd.AddCommand(cmdInstanceList)
	Cmd.AddCommand(cmdInstanceInfo)

}

// cmdCloudInstance ...
var Cmd = &cobra.Command{
	Use:     "instance",
	Short:   "Instance commands: ovhcli cloud instance --help",
	Long:    `Instance commands: ovhcli cloud instance <command>`,
	Aliases: []string{"in"},
}
