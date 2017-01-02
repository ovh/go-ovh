package project

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
	cmdProjectRegion.AddCommand(cmdProjectRegionList)
	cmdProjectRegionList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display cloud region details")
}

var (
	cmdProjectRegion = &cobra.Command{
		Use:   "region",
		Short: "Project region management",
		Run: func(cmd *cobra.Command, args []string) {
			common.WrongUsage(cmd)
		},
	}

	cmdProjectRegionList = &cobra.Command{
		Use:   "list",
		Short: "List all regions: ovhcli cloud region list",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := ovh.NewDefaultClient()
			common.Check(err)

			if projectName != "" {
				p, err := client.CloudProjectInfoByName(projectName)
				common.Check(err)
				projectID = p.ID
			}

			if projectID == "" {
				common.WrongUsage(cmd)
			}

			regions, err := client.CloudListRegions(projectID)
			common.Check(err)

			if withDetails {
				regionsComplete := []ovh.Region{}
				for _, region := range regions {
					r, err := client.CloudInfoRegion(projectID, region.Region)
					common.Check(err)
					regionsComplete = append(regionsComplete, *r)
				}
				regions = regionsComplete

			}

			common.FormatOutputDef(regions)
		},
	}

	cmdProjectRegionByName = &cobra.Command{
		Use:   "info <name>",
		Short: "Get info for region",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				common.WrongUsage(cmd)
			}
			client, err := ovh.NewDefaultClient()
			common.Check(err)

			if projectName != "" {
				p, err := client.CloudProjectInfoByName(projectName)
				common.Check(err)
				projectID = p.ID
			}

			if projectID == "" {
				common.WrongUsage(cmd)
			}

			r, err := client.CloudInfoRegion(projectID, args[0])
			common.Check(err)

			common.FormatOutputDef(r)
		},
	}
)
