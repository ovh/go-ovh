package queue

import (
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas/queue/key"
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas/queue/metrics"
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas/queue/region"
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas/queue/role"
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas/queue/service"
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas/queue/topic"
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas/queue/user"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(key.Cmd)
	Cmd.AddCommand(metrics.Cmd)
	Cmd.AddCommand(region.Cmd)
	Cmd.AddCommand(role.Cmd)
	Cmd.AddCommand(service.Cmd)
	Cmd.AddCommand(topic.Cmd)
	Cmd.AddCommand(user.Cmd)
}

// Cmd cmdCloudQueue ...
var Cmd = &cobra.Command{
	Use:   "queue",
	Short: "Queue commands: ovhcli dbaas queue --help",
	Long:  `Queue commands: ovhcli dbaas queue <command>`,
}
