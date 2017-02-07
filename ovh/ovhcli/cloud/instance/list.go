package instance

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

// var withDetails bool

func init() {
	// cmdInstanceList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display cloud instance details")
	cmdInstanceList.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
}

var cmdInstanceList = &cobra.Command{
	Use:   "list",
	Short: "List all instance: ovhcli cloud instance list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		instances, err := client.CloudListInstance(projectID)
		common.Check(err)

		/*	if withDetails {
			instancesComplete := []sdk.Instance{}
			for _, instance := range instances {
				i, err := client.VrackInfo(instance.Name)
				common.Check(err)
				instancesComplete = append(instancesComplete, *i)
			}
			instances = instancesComplete
		} */

		common.FormatOutputDef(instances)
	},
}
