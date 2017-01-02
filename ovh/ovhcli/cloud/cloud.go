package cloud

import (
	"github.com/runabove/go-sdk/ovh/ovhcli/cloud/instance"
	"github.com/runabove/go-sdk/ovh/ovhcli/cloud/network"
	"github.com/runabove/go-sdk/ovh/ovhcli/cloud/project"
	"github.com/runabove/go-sdk/ovh/ovhcli/cloud/sshkey"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(project.Cmd)
	Cmd.AddCommand(instance.Cmd)
	Cmd.AddCommand(network.Cmd)
	Cmd.AddCommand(sshkey.Cmd)
}

// Cmd project
var (
	Cmd = &cobra.Command{
		Use:     "cloud",
		Short:   "Project commands: ovhcli cloud --help",
		Long:    `Project commands: ovhcli cloud <command>`,
		Aliases: []string{"cl"},
	}
)
