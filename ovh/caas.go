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

// ContainersServicesList ...
func (c *Client) ContainersServicesList() ([]ContainersService, error) {
	var contlist []string
	e := c.Get("/caas/containers", &contlist)
	containersservices := []ContainersService{}
	for _, cont := range contlist {
		containersservices = append(containersservices, ContainersService{Name: cont})
	}
	return containersservices, e
}

// ContainersServiceInfo ...
func (c *Client) ContainersServiceInfo(containerservid string) (*ContainersService, error) {
	containersservice := &ContainersService{}
	e := c.Get(queryEscape("/caas/containers/%s", containerservid), &containersservice)

	return containersservice, e
}
