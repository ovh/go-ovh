// ovhcli offers to manage your Ovh services
package main

import (
	"fmt"
	"os"

	"github.com/runabove/go-sdk/ovh/ovhcli/caas"
	"github.com/runabove/go-sdk/ovh/ovhcli/cloud"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/runabove/go-sdk/ovh/ovhcli/connect"
	"github.com/runabove/go-sdk/ovh/ovhcli/dbaas"
	"github.com/runabove/go-sdk/ovh/ovhcli/domain"
	"github.com/runabove/go-sdk/ovh/ovhcli/order"
	"github.com/runabove/go-sdk/ovh/ovhcli/telephony"
	"github.com/runabove/go-sdk/ovh/ovhcli/version"
	"github.com/runabove/go-sdk/ovh/ovhcli/vrack"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ovhcli",
	Short: "OVH - Command Line Tool",
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&common.Format, "format", "f", "pretty", "choose format output. One of 'json', 'yaml' and 'pretty'")
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "verbose output")

	addCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//AddCommands adds child commands to the root command rootCmd.
func addCommands() {
	rootCmd.AddCommand(caas.Cmd)
	rootCmd.AddCommand(domain.Cmd)
	rootCmd.AddCommand(cloud.Cmd)
	rootCmd.AddCommand(dbaas.Cmd)
	rootCmd.AddCommand(telephony.Cmd)

	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(vrack.Cmd)

	rootCmd.AddCommand(connect.Cmd)

	rootCmd.AddCommand(order.Cmd)

	rootCmd.AddCommand(autocompleteCmd)
}

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete <path>",
	Short: "Generate bash autocompletion file for ovhcli",
	Long:  `Generate bash autocompletion file for ovhcli`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintf(os.Stderr, "Wrong usage: ovhcli autocomplete <path>\n")
			return
		}
		rootCmd.GenBashCompletionFile(args[0])
		fmt.Fprintf(os.Stderr, "Completion file generated.\n")
		fmt.Fprintf(os.Stderr, "You may now run `source %s`\n", args[0])
	},
}
