package sshkey

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdCloudSSHKeyList)
	Cmd.AddCommand(cmdCloudSSHKeyCreate)
	Cmd.AddCommand(cmdCloudSSHKeyDelete)

}

// cmdCloudSSHkey ...
var (
	projectID string
	pubkeyID  string
	name      string

	Cmd = &cobra.Command{
		Use:     "sshkey",
		Short:   "sshkey commands: ovhcli cloud sshkey --help",
		Long:    `Regisshkeyon commands: ovhcli cloud sshkey <command>`,
		Aliases: []string{"ssh"},
	}
)
