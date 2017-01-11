package dbaas

import (
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas/queue"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(queue.Cmd)
}

// Cmd project
var (
	Cmd = &cobra.Command{
		Use:   "dbaas",
		Short: "Project commands: ovhcli dbaas --help",
		Long:  `Project commands: ovhcli dbaas <command>`,
	}
)
