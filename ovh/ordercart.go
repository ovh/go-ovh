package ovh

import (
	"strconv"

	"github.com/runabove/go-sdk/ovh/types"
)

// OrderCartList list all your cart
func (c *Client) OrderCartList() ([]types.OrderCart, error) {
	var ids []string
	e := c.Get("/order/cart", &ids)
	carts := []types.OrderCart{}
	for _, id := range ids {
		carts = append(carts, types.OrderCart{CartID: id})
	}
	return carts, e
}

// OrderCartInfo retrieve all infos of one of your cart
func (c *Client) OrderCartInfo(cartID string) (*types.OrderCart, error) {
	cart := &types.OrderCart{}
	err := c.Get(queryEscape("/order/cart/%s", cartID), cart)
	return cart, err
}

// OrderCreateCart create a new cart
func (c *Client) OrderCreateCart(cartCreateReq types.OrderCartPost) (*types.OrderCart, error) {
	cart := &types.OrderCart{}
	e := c.Post("/order/cart", cartCreateReq, cart)
	return cart, e
}

// OrderUpdateCart update a cart
func (c *Client) OrderUpdateCart(cartID string, cartUpdateReq types.OrderCartPost) (*types.OrderCart, error) {
	cart := &types.OrderCart{}
	e := c.Put(queryEscape("/order/cart/%s", cartID), cartUpdateReq, cart)
	return cart, e
}

// OrderDeleteCart delete a cart
func (c *Client) OrderDeleteCart(cartID string) error {
	e := c.Delete(queryEscape("/order/cart/%s", cartID), nil)
	return e
}

// OrderAssignCart assign to connected user a cart
func (c *Client) OrderAssignCart(cartID string) error {
	e := c.Post(queryEscape("/order/cart/%s/assign", cartID), nil, nil)
	return e
}

// OrderSummaryCart get a summary of your current order
func (c *Client) OrderSummaryCart(cartID string) (*types.Order, error) {
	order := &types.Order{}
	e := c.Get(queryEscape("/order/cart/%s/summary", cartID), order)
	return order, e
}

// OrderGetCheckoutCart get prices and contracts information for your cart
func (c *Client) OrderGetCheckoutCart(cartID string) (*types.Order, error) {
	order := &types.Order{}
	e := c.Get(queryEscape("/order/cart/%s/checkout", cartID), order)
	return order, e
}

// OrderPostCheckoutCart validate your shopping and create order
func (c *Client) OrderPostCheckoutCart(cartID string, waiveRetractationPeriod bool) (*types.Order, error) {
	order := &types.Order{}

	data := struct {
		WaiveRetractationPeriod bool `json:"waiveRetractationPeriod"`
	}{
		waiveRetractationPeriod,
	}

	e := c.Post(queryEscape("/order/cart/%s/checkout", cartID), data, order)
	return order, e
}

// OrderCartItemList list all items in your cart
func (c *Client) OrderCartItemList(cartID string) ([]types.OrderCartItem, error) {
	var ids []int64
	e := c.Get(queryEscape("/order/cart/%s/item", cartID), &ids)
	items := []types.OrderCartItem{}
	for _, id := range ids {
		items = append(items, types.OrderCartItem{ItemID: id})
	}
	return items, e
}

// OrderCartItemInfo retrieve info of a cart item
func (c *Client) OrderCartItemInfo(cartID string, itemID int64) (*types.OrderCartItem, error) {
	item := &types.OrderCartItem{}
	err := c.Get(queryEscape("/order/cart/%s/item/%s", cartID, strconv.FormatInt(itemID, 10)), item)
	return item, err
}

// OrderUpdateCartItem update a cart item
func (c *Client) OrderUpdateCartItem(cartID string, itemID int64, duration string, quantity int) (*types.OrderCartItem, error) {
	item := &types.OrderCartItem{}

	data := struct {
		Duration string `json:"duration,omitempty"`
		Quantity int    `json:"quantity,omitempty"`
	}{
		duration,
		quantity,
	}

	err := c.Put(queryEscape("/order/cart/%s/item/%s", cartID, strconv.FormatInt(itemID, 10)), data, item)
	return item, err
}

// OrderDeleteCartItem delete a cart item
func (c *Client) OrderDeleteCartItem(cartID string, itemID int64) (*types.OrderCartItem, error) {
	err := c.Delete(queryEscape("/order/cart/%s/item/%s", cartID, strconv.FormatInt(itemID, 10)), nil)
	return nil, err
}

// OrderCartConfigurationsList list all configurations for an item
func (c *Client) OrderCartConfigurationsList(cartID string, itemID int64) ([]types.OrderCartConfigurationItem, error) {
	var ids []int64
	e := c.Get(queryEscape("/order/cart/%s/item/%s/configuration", cartID, strconv.FormatInt(itemID, 10)), &ids)
	configs := []types.OrderCartConfigurationItem{}
	for _, id := range ids {
		configs = append(configs, types.OrderCartConfigurationItem{ID: id})
	}
	return configs, e
}

// OrderCartConfigurationInfo get a configuration for an item
func (c *Client) OrderCartConfigurationInfo(cartID string, itemID int64, configID int64) (*types.OrderCartConfigurationItem, error) {
	config := &types.OrderCartConfigurationItem{}
	err := c.Get(queryEscape("/order/cart/%s/item/%s/configuration/%s", cartID, strconv.FormatInt(itemID, 10), strconv.FormatInt(configID, 10)), config)
	return config, err
}

// OrderCartAddConfiguration add a configuration on an item
func (c *Client) OrderCartAddConfiguration(cartID string, itemID int64, label string, value string) (*types.OrderCartItem, error) {
	item := &types.OrderCartItem{}

	data := struct {
		Label string `json:"label,omitempty"`
		Value string `json:"value,omitempty"`
	}{
		label,
		value,
	}
	err := c.Post(queryEscape("/order/cart/%s/item/%s/configuration", cartID, strconv.FormatInt(itemID, 10)), data, item)
	return item, err
}

// OrderCartDeleteConfiguration remove a configuration from an item
func (c *Client) OrderCartDeleteConfiguration(cartID string, itemID int64, configID int64) (*types.OrderCartItem, error) {
	err := c.Delete(queryEscape("/order/cart/%s/item/%s/configuration/%s", cartID, strconv.FormatInt(itemID, 10), strconv.FormatInt(configID, 10)), nil)
	return nil, err
}

// OrderCartRequiredConfigurations get required configurations for an item
func (c *Client) OrderCartRequiredConfigurations(cartID string, itemID int64) ([]types.OrderCartConfigurationRequirements, error) {
	var configs []types.OrderCartConfigurationRequirements
	e := c.Get(queryEscape("/order/cart/%s/item/%s/requiredConfiguration", cartID, strconv.FormatInt(itemID, 10)), &configs)
	return configs, e
}
