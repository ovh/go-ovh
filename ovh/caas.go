package ovh

import (
	"github.com/runabove/go-sdk/ovh/types"
)

// ContainersServicesList list all your containers
func (c *Client) ContainersServicesList(withDetails bool) ([]types.DockerStack, error) {
	var names []string
	if err := c.Get("/caas/containers", &names); err != nil {
		return nil, err
	}

	containers := []types.DockerStack{}
	for _, name := range names {
		containers = append(containers, types.DockerStack{Name: name})
	}

	if !withDetails {
		return containers, nil
	}

	containersChan, errChan := make(chan types.DockerStack), make(chan error)
	for _, container := range containers {
		go func(container types.DockerStack) {
			d, err := c.ContainersServiceInfo(container.Name)
			if err != nil {
				errChan <- err
				return
			}
			containersChan <- *d
		}(container)
	}

	containersComplete := []types.DockerStack{}

	for i := 0; i < len(containers); i++ {
		select {
		case containers := <-containersChan:
			containersComplete = append(containersComplete, containers)
		case err := <-errChan:
			return nil, err
		}
	}

	return containersComplete, nil
}

// ContainersServiceInfo retrieve all infos of one of your containers
func (c *Client) ContainersServiceInfo(containersName string) (*types.DockerStack, error) {
	containers := &types.DockerStack{}
	err := c.Get(queryEscape("/caas/containers/%s", containersName), containers)
	return containers, err
}
