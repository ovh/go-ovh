package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION of ovhcli
const VERSION = "0.1"

// Cmd version
var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Display Version of ovhcli: ovhcli version",
	Long:    `ovhcli version`,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("ovhcli version:", VERSION)

	},
}
