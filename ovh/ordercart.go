package ovh

import (
	"strconv"
	"time"
)

// OrderCartCreateReq defines the fields for a Cart creation request
type OrderCartCreateReq struct {
	Expire        *time.Time `json:"expire,omitempty"`
	Description   string     `json:"description,omitempty"`
	OVHSubsidiary string     `json:"ovhSubsidiary,omitempty"`
}

// OrderCartUpdateReq defines the fields for a Cart update request
type OrderCartUpdateReq struct {
	Expire      *time.Time `json:"expire,omitempty"`
	Description string     `json:"description,omitempty"`
}

// OrderCartList list all your cart
func (c *Client) OrderCartList() ([]OrderCart, error) {
	var ids []string
	e := c.Get("/order/cart", &ids)
	carts := []OrderCart{}
	for _, id := range ids {
		carts = append(carts, OrderCart{CartID: id})
	}
	return carts, e
}

// OrderCartInfo retrieve all infos of one of your cart
func (c *Client) OrderCartInfo(cartID string) (*OrderCart, error) {
	cart := &OrderCart{}
	err := c.Get(queryEscape("/order/cart/%s", cartID), cart)
	return cart, err
}

// OrderCreateCart create a new cart
func (c *Client) OrderCreateCart(cartCreateReq OrderCartCreateReq) (*OrderCart, error) {
	cart := &OrderCart{}
	e := c.Post("/order/cart", cartCreateReq, cart)
	return cart, e
}

// OrderUpdateCart update a cart
func (c *Client) OrderUpdateCart(cartID string, cartUpdateReq OrderCartUpdateReq) (*OrderCart, error) {
	cart := &OrderCart{}
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
func (c *Client) OrderSummaryCart(cartID string) (*Order, error) {
	order := &Order{}
	e := c.Get(queryEscape("/order/cart/%s/summary", cartID), order)
	return order, e
}

// OrderGetCheckoutCart get prices and contracts information for your cart
func (c *Client) OrderGetCheckoutCart(cartID string) (*Order, error) {
	order := &Order{}
	e := c.Get(queryEscape("/order/cart/%s/checkout", cartID), order)
	return order, e
}

// OrderPostCheckoutCart validate your shopping and create order
func (c *Client) OrderPostCheckoutCart(cartID string, waiveRetractationPeriod bool) (*Order, error) {
	order := &Order{}

	data := struct {
		WaiveRetractationPeriod bool `json:"waiveRetractationPeriod"`
	}{
		waiveRetractationPeriod,
	}

	e := c.Post(queryEscape("/order/cart/%s/checkout", cartID), data, order)
	return order, e
}

// OrderCartItemList list all items in your cart
func (c *Client) OrderCartItemList(cartID string) ([]OrderCartItem, error) {
	var ids []int
	e := c.Get(queryEscape("/order/cart/%s/item", cartID), &ids)
	items := []OrderCartItem{}
	for _, id := range ids {
		items = append(items, OrderCartItem{ItemID: id})
	}
	return items, e
}

// OrderCartItemInfo retrieve info of a cart item
func (c *Client) OrderCartItemInfo(cartID string, itemID int) (*OrderCartItem, error) {
	item := &OrderCartItem{}
	err := c.Get(queryEscape("/order/cart/%s/item/%s", cartID, strconv.Itoa(itemID)), item)
	return item, err
}

// OrderUpdateCartItem update a cart item
func (c *Client) OrderUpdateCartItem(cartID string, itemID int, duration string, quantity int) (*OrderCartItem, error) {
	item := &OrderCartItem{}

	data := struct {
		Duration string `json:"duration,omitempty"`
		Quantity int    `json:"quantity,omitempty"`
	}{
		duration,
		quantity,
	}

	err := c.Put(queryEscape("/order/cart/%s/item/%s", cartID, strconv.Itoa(itemID)), data, item)
	return item, err
}

// OrderDeleteCartItem delete a cart item
func (c *Client) OrderDeleteCartItem(cartID string, itemID int) (*OrderCartItem, error) {
	err := c.Delete(queryEscape("/order/cart/%s/item/%s", cartID, strconv.Itoa(itemID)), nil)
	return nil, err
}

// OrderCartConfigurationsList list all configurations for an item
func (c *Client) OrderCartConfigurationsList(cartID string, itemID int) ([]OrderCartConfigurationItem, error) {
	var ids []int
	e := c.Get(queryEscape("/order/cart/%s/item/%s/configuration", cartID, strconv.Itoa(itemID)), &ids)
	configs := []OrderCartConfigurationItem{}
	for _, id := range ids {
		configs = append(configs, OrderCartConfigurationItem{ID: id})
	}
	return configs, e
}

// OrderCartConfigurationInfo get a configuration for an item
func (c *Client) OrderCartConfigurationInfo(cartID string, itemID int, configID int) (*OrderCartConfigurationItem, error) {
	config := &OrderCartConfigurationItem{}
	err := c.Get(queryEscape("/order/cart/%s/item/%s/configuration/%s", cartID, strconv.Itoa(itemID), strconv.Itoa(configID)), config)
	return config, err
}

// OrderCartAddConfiguration add a configuration on an item
func (c *Client) OrderCartAddConfiguration(cartID string, itemID int, label string, value string) (*OrderCartItem, error) {
	item := &OrderCartItem{}

	data := struct {
		Label string `json:"label,omitempty"`
		Value string `json:"value,omitempty"`
	}{
		label,
		value,
	}
	err := c.Post(queryEscape("/order/cart/%s/item/%s/configuration", cartID, strconv.Itoa(itemID)), data, item)
	return item, err
}

// OrderCartDeleteConfiguration remove a configuration from an item
func (c *Client) OrderCartDeleteConfiguration(cartID string, itemID int, configID int) (*OrderCartItem, error) {
	err := c.Delete(queryEscape("/order/cart/%s/item/%s/configuration/%s", cartID, strconv.Itoa(itemID), strconv.Itoa(configID)), nil)
	return nil, err
}

// OrderCartRequiredConfigurations get required configurations for an item
func (c *Client) OrderCartRequiredConfigurations(cartID string, itemID int) ([]OrderCartConfigurationRequirements, error) {
	var configs []OrderCartConfigurationRequirements
	e := c.Get(queryEscape("/order/cart/%s/item/%s/requiredConfiguration", cartID, strconv.Itoa(itemID)), &configs)
	return configs, e
}
