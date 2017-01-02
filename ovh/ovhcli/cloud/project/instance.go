package project

import (
	"fmt"

	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/spf13/cobra"
)

func init() {
	cmdProjectInstance.AddCommand(cmdProjectInstanceList)
	cmdProjectInstance.AddCommand(cmdProjectInstanceCreate)

	cmdProjectInstanceCreate.Flags().StringVar(&instanceImage, "image", "", "Define image")
	cmdProjectInstanceCreate.Flags().StringVar(&instanceFlavor, "flavor", "", "Define flavor")
	cmdProjectInstanceCreate.Flags().StringVar(&instanceSSHKey, "sshKey", "", "Define ssh key")
	cmdProjectInstanceCreate.MarkFlagRequired("image")
	cmdProjectInstanceCreate.MarkFlagRequired("flavor")
	cmdProjectInstanceCreate.MarkFlagRequired("sshKey")
}

var (
	instanceImage  string
	instanceFlavor string
	instanceSSHKey string

	cmdProjectInstance = &cobra.Command{
		Use:   "instance",
		Short: "Project instances management",
		Run: func(cmd *cobra.Command, args []string) {
			common.WrongUsage(cmd)
		},
	}

	cmdProjectInstanceCreate = &cobra.Command{
		Use:   "create <Server>",
		Short: "Create instance",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				common.WrongUsage(cmd)
			}

			client, err := ovh.NewDefaultClient()
			common.Check(err)

			if projectName != "" {
				p, err := client.CloudProjectInfoByName(projectName)
				common.Check(err)
				projectID = p.ID
			}

			if projectID == "" || regionName == "" {
				common.WrongUsage(cmd)
			}

			imgs, err := client.CloudProjectImagesList(projectID, regionName)
			common.Check(err)

			snaps, err := client.CloudProjectSnapshotsList(projectID, regionName)
			common.Check(err)

			imgs = append(imgs, snaps...)

			var img *ovh.Image
			for i := range imgs {
				if imgs[i].Name == instanceImage && imgs[i].Region == regionName {
					img = &imgs[i]
					break
				}
			}

			if img == nil {
				common.Check(fmt.Errorf("Image %s not found", instanceImage))
			}

			flavors, err := client.CloudProjectFlavorsList(projectID, regionName)
			common.Check(err)

			var f *ovh.Flavor
			for i := range flavors {
				if flavors[i].Name == instanceFlavor && flavors[i].Region == regionName {
					f = &flavors[i]
					break
				}
			}

			if f == nil {
				common.Check(fmt.Errorf("Flavor %s not found", instanceFlavor))
			}

			sshkeys, err := client.CloudProjectSSHKeyList(projectID)
			common.Check(err)

			var k *ovh.Sshkey
			for i := range sshkeys {
				if sshkeys[i].Name == instanceSSHKey {
					k = &sshkeys[i]
					break
				}
			}

			if k == nil {
				common.Check(fmt.Errorf("SSH Key %s not found", instanceSSHKey))
			}

			ins, err := client.CloudCreateInstance(projectID, args[0], k.ID, f.ID, img.ID, regionName)
			common.Check(err)

			common.FormatOutputDef(ins)
		},
	}

	cmdProjectInstanceList = &cobra.Command{}
)
