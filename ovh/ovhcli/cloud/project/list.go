package project

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/runabove/go-sdk/ovh/types"

	"github.com/spf13/cobra"
)

func init() {
	cmdProjectList.Flags().BoolVarP(&withDetails, "withDetails", "", false, "Display project details")
}

var cmdProjectList = &cobra.Command{
	Use:   "list",
	Short: "List all projects: ovhcli project list",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		common.Check(err)

		projects, err := client.CloudProjectsList()

		if withDetails {
			projectsComplete := []types.CloudProject{}
			for _, project := range projects {
				p, e := client.CloudProjectInfoByID(project.ProjectID)
				common.Check(e)
				projectsComplete = append(projectsComplete, *p)
			}
			projects = projectsComplete
		}

		common.Check(err)
		common.FormatOutputDef(projects)
	},
}
