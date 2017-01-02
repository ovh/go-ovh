package sshkey

import (
	"fmt"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
	cmdCloudSSHKeyDelete.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
	cmdCloudSSHKeyDelete.PersistentFlags().StringVarP(&pubkeyID, "sshkeyID", "", "", "Your sshkey ID to delete")
}

var cmdCloudSSHKeyDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete Cloud SSH key: ovhcli cloud sshkey delete",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		err = client.CloudProjectSSHKeyDelete(projectID, pubkeyID)
		common.Check(err)

		fmt.Printf("Public SSH key %s deleted:\n", pubkeyID)

	},
}
