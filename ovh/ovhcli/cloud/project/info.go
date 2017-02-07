package project

import (
	"github.com/spf13/cobra"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/runabove/go-sdk/ovh/types"
)

var (
	cmdProjectInfo = &cobra.Command{
		Use:   "info",
		Short: "Info about a project",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := ovh.NewDefaultClient()
			common.Check(err)

			if projectID == "" && projectName == "" {
				common.WrongUsage(cmd)
			}
			var p *types.CloudProject
			if projectID != "" {
				p, err = client.CloudProjectInfoByID(projectID)
			} else {
				p, err = client.CloudProjectInfoByName(projectName)
			}
			common.Check(err)
			common.FormatOutputDef(p)
		},
	}
)
