package project

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/spf13/cobra"
)

func init() {
	cmdProjectImage.AddCommand(cmdProjectImageList)
	cmdProjectImage.AddCommand(cmdProjectImageSearch)
}

var (
	cmdProjectImage = &cobra.Command{
		Use:   "image",
		Short: "Project image & snapshots management",
		Run: func(cmd *cobra.Command, args []string) {
			common.WrongUsage(cmd)
		},
	}

	cmdProjectImageList = &cobra.Command{
		Use:   "list",
		Short: "List images & snapshots",
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

			images, err := client.CloudProjectImagesList(projectID, regionName)
			common.Check(err)

			snapshots, err := client.CloudProjectSnapshotsList(projectID, regionName)
			common.Check(err)

			images = append(images, snapshots...)

			common.FormatOutputDef(images)
		},
	}

	cmdProjectImageSearch = &cobra.Command{
		Use:   "search <term0> [term1] [...] [termN]",
		Short: "Search images & snapshots",
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

			images, err := client.CloudProjectImagesSearch(projectID, regionName, args...)
			common.Check(err)
			common.FormatOutputDef(images)
		},
	}
)
