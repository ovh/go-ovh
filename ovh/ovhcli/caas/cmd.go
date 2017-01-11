package caas

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdContainersServicesList)
	Cmd.AddCommand(cmdContainersServiceInfo)

}

// Cmd project
var Cmd = &cobra.Command{
	Use:     "caas",
	Short:   "CaaS commands: ovhcli caas --help",
	Long:    `CaaS commands: ovhcli caas <command>`,
	Aliases: []string{"c"},
}
