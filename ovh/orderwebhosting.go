package ovh

import (
	"errors"

	"github.com/runabove/go-sdk/ovh/types"
)

// OrderAddProductWebHosting post a new webHosting product in your cart
func (c *Client) OrderAddProductWebHosting(cartID string, orderCartWebHostingPost types.OrderCartWebHostingPost) (*types.OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	item := &types.OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/webHosting", cartID), orderCartWebHostingPost, item)
	return item, err
}
