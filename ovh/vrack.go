package ovh

import (
	"github.com/runabove/go-sdk/ovh/types"
)

// VrackList ...
func (c *Client) VrackList() ([]types.Vrack, error) {
	ids := []string{}
	e := c.Get("/vrack", &ids)
	vracks := []types.Vrack{}
	for _, id := range ids {
		vracks = append(vracks, types.Vrack{Name: id})
	}
	return vracks, e
}

// VrackInfo ...
func (c *Client) VrackInfo(vrackName string) (*types.Vrack, error) {
	vrack := &types.Vrack{}
	err := c.Get(queryEscape("/vrack/%s", vrackName), vrack)
	return vrack, err
}
