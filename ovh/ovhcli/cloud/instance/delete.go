package instance

import (
	"fmt"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var instanceID string

func init() {
	cmdInstanceDelete.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
	cmdInstanceDelete.PersistentFlags().StringVarP(&instanceID, "instanceID", "", "", "Your Instance ID to delete")

}

var cmdInstanceDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete Cloud Public Instance: ovhcli cloud instance delete",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		err = client.CloudDeleteInstance(projectID, instanceID)
		common.Check(err)

		fmt.Printf("Instance %s deleted:\n", instanceID)

	},
}
