package ovh

// Vrack ...
type Vrack struct {

	// "Vrack name"
	Name string `json:"name,omitempty"`

	// "Vrack decription"
	Description string `json:"description,omitempty"`
}

// VrackList ...
func (c *Client) VrackList() ([]Vrack, error) {
	ids := []string{}
	e := c.Get("/vrack", &ids)
	vracks := []Vrack{}
	for _, id := range ids {
		vracks = append(vracks, Vrack{Name: id})
	}
	return vracks, e
}

// VrackInfo ...
func (c *Client) VrackInfo(vrackName string) (*Vrack, error) {
	vrack := &Vrack{}
	err := c.Get(queryEscape("/vrack/%s", vrackName), vrack)
	return vrack, err
}
