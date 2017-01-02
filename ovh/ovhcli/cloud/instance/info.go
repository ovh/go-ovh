package instance

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
	cmdInstanceInfo.PersistentFlags().StringVarP(&instanceID, "instanceID", "", "", "Your Instance ID")
	cmdInstanceInfo.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
}

var cmdInstanceInfo = &cobra.Command{
	Use:   "info",
	Short: "Info about an cloud instance: ovhcli cloud instance info",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		instance, err := client.CloudInfoInstance(projectID, instanceID)
		common.Check(err)
		common.FormatOutputDef(instance)
	},
}
