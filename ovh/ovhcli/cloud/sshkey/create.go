package sshkey

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
	cmdCloudSSHKeyCreate.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
	cmdCloudSSHKeyCreate.PersistentFlags().StringVarP(&pubkeyID, "pubkeyID", "", "", "Your sshkey ID to put")
	cmdCloudSSHKeyCreate.PersistentFlags().StringVarP(&name, "name", "", "", "Your sshkey name to put")
}

var cmdCloudSSHKeyCreate = &cobra.Command{
	Use:   "create",
	Short: "Create Cloud ssh key: ovhcli cloud sshkey create",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		s, err := client.CloudProjectSSHKeyCreate(projectID, pubkeyID, name)
		common.Check(err)
		common.FormatOutputDef(s)
	},
}
