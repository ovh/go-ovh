package domain

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdDomainList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display domain details")
}

var cmdDomainList = &cobra.Command{
	Use:   "list",
	Short: "List all domains: ovhcli domain list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		domains, err := client.DomainList(withDetails)
		common.Check(err)

		common.FormatOutputDef(domains)
	},
}