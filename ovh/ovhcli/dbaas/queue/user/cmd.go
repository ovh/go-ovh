package user

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var id string
var name string
var userName string
var userID string

func init() {
	Cmd.AddCommand(cmdChangePassword)
	Cmd.AddCommand(cmdInfo)
	Cmd.AddCommand(cmdList)

	Cmd.PersistentFlags().StringVarP(&id, "id", "", "", "Your Application ID")
	Cmd.PersistentFlags().StringVarP(&name, "name", "", "", "Your Application Name")
	Cmd.PersistentFlags().StringVarP(&userName, "user", "", "", "User Name")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "user",
	Short: "Queue user commands: ovhcli dbaas queue user --help",
	Long:  `Queue user commands: ovhcli dbaas queue user <command>`,
}

func checkUser() {
	if userID == "" {
		fmt.Fprintf(os.Stderr, "User %s not found\n", userName)
		os.Exit(1)
	}
}
