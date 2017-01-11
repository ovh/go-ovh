package domain

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdDomainInfo = &cobra.Command{
	Use:   "info <domain>",
	Short: "Retrieve domain info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		domain := args[0]

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		d, err := client.DomainInfo(domain)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}
