package ovh

import (
	"fmt"
	"strings"

	"github.com/runabove/go-sdk/ovh/types"
)

const (
	// CustomerInterface is the URL of the customer interface, for error messages
	CustomerInterface = "https://www.ovh.com/manager/cloud/index.html"
)

// CloudProjectsList returns a list of string project ID
func (c *Client) CloudProjectsList() ([]types.CloudProject, error) {
	projects := []types.CloudProject{}
	ids := []string{}
	if err := c.Get("/cloud/project", &ids); err != nil {
		return nil, err
	}
	for _, id := range ids {
		projects = append(projects, types.CloudProject{ProjectID: id})
	}
	return projects, nil
}

// CloudProjectInfoByID return the details of a project given a project id
func (c *Client) CloudProjectInfoByID(projectID string) (*types.CloudProject, error) {
	project := &types.CloudProject{}
	err := c.Get(queryEscape("/cloud/project/%s", projectID), &project)
	return project, err
}

// CloudProjectInfoByName returns the details of a project given its name.
func (c *Client) CloudProjectInfoByName(projectDescription string) (project *types.CloudProject, err error) {
	// get project list
	projects, err := c.CloudProjectsList()
	if err != nil {
		return nil, err
	}

	// If projectDescription is a valid projectID return it.
	for _, p := range projects {
		if p.ProjectID == projectDescription {
			return c.CloudProjectInfoByID(p.ProjectID)
		}
	}

	// Attempt to find a project matching projectDescription. This is potentially slow
	for _, p := range projects {
		project, err := c.CloudProjectInfoByID(p.ProjectID)
		if err != nil {
			return nil, err
		}

		if project.Description == projectDescription {
			return project, nil
		}
	}

	// Ooops
	return nil, fmt.Errorf("Project '%s' does not exist on OVH cloud. To create or rename a project, please visit %s", projectDescription, CustomerInterface)
}

// CloudListRegions return a list of network regions
func (c *Client) CloudListRegions(projectID string) ([]types.CloudRegionDetail, error) {
	var resultsreq []string
	if err := c.Get(queryEscape("/cloud/project/%s/region", projectID), &resultsreq); err != nil {
		return nil, err
	}
	regions := []types.CloudRegionDetail{}
	for _, resultreq := range resultsreq {
		regions = append(regions, types.CloudRegionDetail{Name: resultreq})
	}
	return regions, nil
}

// CloudInfoRegion return services status on a region
func (c *Client) CloudInfoRegion(projectID, regionName string) (*types.CloudRegionDetail, error) {
	region := &types.CloudRegionDetail{}
	err := c.Get(queryEscape("/cloud/project/%s/region/%s", projectID, regionName), region)
	return region, err
}

// CloudGetInstance finds a VM instance given a name or an ID
func (c *Client) CloudGetInstance(projectID, instanceID string) (instance *types.CloudInstance, err error) {
	err = c.Get(queryEscape("/cloud/project/%s/instance/%s", projectID, instanceID), &instance)
	return instance, nil
}

// CloudCreateInstance start a new public cloud instance and returns resulting object
func (c *Client) CloudCreateInstance(projectID, name, pubkeyID, flavorID, imageID, region string) (instance *types.CloudInstance, err error) {
	instanceReq := types.CloudInstance{
		Name:     name,
		SSHKeyID: pubkeyID,
		FlavorID: flavorID,
		ImageID:  imageID,
		Region:   region,
	}
	err = c.Post(queryEscape("/cloud/project/%s/instance", projectID), instanceReq, &instance)
	return instance, err
}

// CloudDeleteInstance stops and destroys a public cloud instance
func (c *Client) CloudDeleteInstance(projectID, instanceID string) error {
	err := c.Delete(queryEscape("/cloud/project/%s/instance/%s", projectID, instanceID), nil)
	if apierror, ok := err.(*APIError); ok && apierror.Code == 404 {
		err = nil
	}
	return err
}

// CloudListInstance show cloud instance(s)
func (c *Client) CloudListInstance(projectID string) ([]types.CloudInstance, error) {
	instances := []types.CloudInstance{}
	err := c.Get(queryEscape("/cloud/project/%s/instance", projectID), &instances)
	return instances, err
}

// CloudInfoInstance give info about cloud instance
func (c *Client) CloudInfoInstance(projectID, instanceID string) (*types.CloudInstance, error) {
	instances := &types.CloudInstance{}
	err := c.Get(queryEscape("/cloud/project/%s/instance/%s", projectID, instanceID), &instances)
	return instances, err
}

// CloudInfoNetworkPublic return the list of a public network by given a project id
func (c *Client) CloudInfoNetworkPublic(projectID string) ([]types.CloudNetwork, error) {
	network := []types.CloudNetwork{}
	err := c.Get(queryEscape("/cloud/project/%s/network/public", projectID), &network)
	return network, err
}

// CloudInfoNetworkPrivate return the list of a private network by given a project id
func (c *Client) CloudInfoNetworkPrivate(projectID string) ([]types.CloudNetwork, error) {
	network := []types.CloudNetwork{}
	err := c.Get(queryEscape("/cloud/project/%s/network/private", projectID), &network)
	return network, err
}

// CloudCreateNetworkPrivate create a private network in a vrack
//func (c *Client) CloudCreateNetworkPrivate(projectID, name string, regions []types.CloudRegionDetail, vlanid int) (net *types.CloudNetwork, err error) {
func (c *Client) CloudCreateNetworkPrivate(projectID, name, regions string, vlanid int64) (net *types.CloudNetwork, err error) {

	var project types.CloudProject
	project.ProjectID = projectID
	var network types.CloudNetwork
	network.Name = name
	network.VlanID = vlanid
	//network.[]types.CloudRegionDetail = regions
	err = c.Post(queryEscape("/cloud/project/%s/network/private", projectID), network, &net)
	return net, err
}

// CloudProjectUsersList return the list of users by given a project id
func (c *Client) CloudProjectUsersList(projectID string) ([]types.CloudUser, error) {
	users := []types.CloudUser{}
	return users, c.Get(queryEscape("/cloud/project/%s/user", projectID), &users)
}

// CloudProjectUserCreate return the list of users by given a project id
func (c *Client) CloudProjectUserCreate(projectID, description string) (types.CloudUser, error) {
	data := map[string]string{
		"description": description,
	}
	user := types.CloudUser{}
	return user, c.Post(queryEscape("/cloud/project/%s/user", projectID), data, &user)
}

// CloudProjectRegionList return the region by given a project id
func (c *Client) CloudProjectRegionList(projectID string) ([]string, error) {
	var r []string
	return r, c.Get(queryEscape("/cloud/project/%s/region", projectID), &r)
}

// CloudProjectSSHKeyList return the list of ssh keys by given a project id
func (c *Client) CloudProjectSSHKeyList(projectID string) ([]types.CloudSSHKey, error) {
	sshkeys := []types.CloudSSHKey{}
	return sshkeys, c.Get(queryEscape("/cloud/project/%s/sshkey", projectID), &sshkeys)
}

// CloudProjectSSHKeyInfo return info about a ssh keys
func (c *Client) CloudProjectSSHKeyInfo(projectID, sshkeyID string) (*types.CloudSSHKey, error) {
	sshkeys := &types.CloudSSHKey{}
	return sshkeys, c.Get(queryEscape("/cloud/project/%s/sshkey/%s", projectID, sshkeyID), &sshkeys)
}

// CloudProjectSSHKeyDelete delete a ssh key
func (c *Client) CloudProjectSSHKeyDelete(projectID, sshkeyID string) error {
	err := c.Delete(queryEscape("/cloud/project/%s/sshkey/%s", projectID, sshkeyID), nil)
	if apierror, ok := err.(*APIError); ok && apierror.Code == 404 {
		err = nil
	}
	return err
}

// CloudProjectSSHKeyCreate return the list of users by given a project id
func (c *Client) CloudProjectSSHKeyCreate(projectID, publicKey, name string) (types.CloudSSHKey, error) {
	data := map[string]string{
		"publicKey": publicKey,
		"name":      name,
	}
	sshkey := types.CloudSSHKey{}
	return sshkey, c.Post(queryEscape("/cloud/project/%s/sshkey", projectID), data, &sshkey)
}

//CloudProjectImagesList returns the list of images by given a project id
func (c *Client) CloudProjectImagesList(projectID, region string) ([]types.CloudImage, error) {
	var path string
	if region == "" {
		path = queryEscape("/cloud/project/%s/image", projectID)
	} else {
		path = queryEscape("/cloud/project/%s/image?region=%s", projectID, region)
	}
	images := []types.CloudImage{}
	return images, c.Get(path, &images)
}

//CloudProjectImagesSearch returns the list of images matching terms
func (c *Client) CloudProjectImagesSearch(projectID string, region string, terms ...string) ([]types.CloudImage, error) {
	images, err := c.CloudProjectImagesList(projectID, region)
	if err != nil {
		return nil, err
	}
	snapshots, err := c.CloudProjectSnapshotsList(projectID, region)
	if err != nil {
		return nil, err
	}

	images = append(images, snapshots...)

	res := []types.CloudImage{}
	for i, img := range images {
		for _, t := range terms {
			if strings.Contains(img.ID, t) {
				res = append(res, images[i])
				continue
			}
			if strings.Contains(img.Name, t) {
				res = append(res, images[i])
				continue
			}
			if strings.Contains(img.TType, t) {
				res = append(res, images[i])
				continue
			}
		}
	}
	return res, nil
}

//CloudProjectSnapshotsList returns the list of snapshots by given a project id
func (c *Client) CloudProjectSnapshotsList(projectID, region string) ([]types.CloudImage, error) {
	var path string
	if region == "" {
		path = queryEscape("/cloud/project/%s/snapshot", projectID)

	} else {
		path = queryEscape("/cloud/project/%s/snapshot?region=%s", projectID, region)
	}
	images := []types.CloudImage{}
	return images, c.Get(path, &images)
}

//CloudProjectFlavorsList returns the list of flavors by given a project id
func (c *Client) CloudProjectFlavorsList(projectID, region string) ([]types.CloudFlavor, error) {
	var path string
	if region == "" {
		path = queryEscape("/cloud/project/%s/flavor", projectID)

	} else {
		path = queryEscape("/cloud/project/%s/flavor?region=%s", projectID, region)
	}
	f := []types.CloudFlavor{}
	return f, c.Get(path, &f)
}
