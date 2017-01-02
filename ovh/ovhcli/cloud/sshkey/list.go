package sshkey

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/spf13/cobra"
)

func init() {
	cmdCloudSSHKeyList.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
}

var cmdCloudSSHKeyList = &cobra.Command{
	Use:   "list",
	Short: "List all ssk keys: ovhcli cloud sshkey list",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		sshkeys, err := client.CloudProjectSSHKeyList(projectID)
		common.Check(err)

		sshkeysComplete := []ovh.Sshkey{}
		for _, sshkey := range sshkeys {
			s, err := client.CloudProjectSSHKeyInfo(projectID, sshkey.ID)
			common.Check(err)
			sshkeysComplete = append(sshkeysComplete, *s)
		}

		sshkeys = sshkeysComplete

		common.Check(err)
		common.FormatOutputDef(sshkeys)
	},
}
