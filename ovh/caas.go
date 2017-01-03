package ovh

// ContainersService is a representation of a Containers Service
type ContainersService struct {
	Cluster      string   `json:"cluster,omitempty"`
	CreatedAt    string   `json:"createdAt,omitempty"`
	Frameworks   []string `json:"frameworks,omitempty"`
	LoadBalancer string   `json:"loadBalancer,omitempty"`
	Metrics      *struct {
		Resources *struct {
			CPU int `json:"cpu,omitempty"`
			Mem int `json:"mem,omitempty"`
		} `json:"resources,omitempty"`
		UsedResources *struct {
			CPU float64 `json:"cpu,omitempty"`
			Mem int     `json:"mem,omitempty"`
		} `json:"usedResources,omitempty"`
	} `json:"metrics,omitempty"`
	Name      string   `json:"name,omitempty"`
	Slaves    []string `json:"slaves,omitempty"`
	State     string   `json:"state,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
}

// ContainersServicesList list all your containers
func (c *Client) ContainersServicesList(withDetails bool) ([]ContainersService, error) {
	var names []string
	if err := c.Get("/caas/containers", &names); err != nil {
		return nil, err
	}

	containers := []ContainersService{}
	for _, name := range names {
		containers = append(containers, ContainersService{Name: name})
	}

	if !withDetails {
		return containers, nil
	}

	containersChan, errChan := make(chan ContainersService), make(chan error)
	for _, container := range containers {
		go func(container ContainersService) {
			d, err := c.ContainersServiceInfo(container.Name)
			if err != nil {
				errChan <- err
				return
			}
			containersChan <- *d
		}(container)
	}

	containersComplete := []ContainersService{}

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
func (c *Client) ContainersServiceInfo(containersName string) (*ContainersService, error) {
	containers := &ContainersService{}
	err := c.Get(queryEscape("/caas/containers/%s", containersName), containers)
	return containers, err
}
