package ovh

import (
	"errors"
)

// OrderPostDomainReq defines the fields for a Cart creation request
type OrderPostDomainReq struct {
	Domain   string `json:"domain,omitempty"`
	Duration string `json:"duration,omitempty"`
	OfferID  string `json:"offerId,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

// OrderPostDomainPacksReq defines the fields for a Cart creation request
type OrderPostDomainPacksReq struct {
	Domain      string `json:"domain,omitempty"`
	Duration    string `json:"duration,omitempty"`
	PlanCode    string `json:"planCode,omitempty"`
	PricingMode string `json:"pricingMode,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
}

// OrderPostDomainOptionReq defines the fields for a Cart creation request
type OrderPostDomainOptionReq struct {
	Duration    string `json:"duration,omitempty"`
	ItemID      int    `json:"itemId,omitempty"`
	PlanCode    string `json:"planCode,omitempty"`
	PricingMode string `json:"pricingMode,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
}

// OrderGetProductsDomain get products about a domain name
func (c *Client) OrderGetProductsDomain(cartID string, domain string) ([]OrderCartProductInformation, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	products := []OrderCartProductInformation{}
	err := c.Get(queryEscape("/order/cart/%s/domain?domain=%s", cartID, domain), &products)
	return products, err
}

// OrderAddProductDomain post a new domain in your cart
func (c *Client) OrderAddProductDomain(cartID string, orderPostDomainReq OrderPostDomainReq) (*OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	domainItem := &OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domain", cartID), orderPostDomainReq, domainItem)
	return domainItem, err
}

// OrderGetProductDomainOptions get informations about a domain name options
func (c *Client) OrderGetProductDomainOptions(cartID string, domain string) ([]OrderCartGenericOptionDefinition, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	options := []OrderCartGenericOptionDefinition{}
	err := c.Get(queryEscape("/order/cart/%s/domain/options?domain=%s", cartID, domain), options)

	return options, err
}

// OrderAddProductDomainOption post an option on a domain item
func (c *Client) OrderAddProductDomainOption(cartID string, orderPostDomainOptionReq OrderPostDomainOptionReq) (*OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	optionItem := &OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domain/options", cartID), orderPostDomainOptionReq, optionItem)
	return optionItem, err
}

// OrderGetProductDomainTransfer get informations about a domain transfer
func (c *Client) OrderGetProductDomainTransfer(cartID string, domain string) ([]OrderCartProductInformation, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	products := []OrderCartProductInformation{}
	err := c.Get(queryEscape("/order/cart/%s/domainTransfer?domain=%s", cartID, domain), products)
	return products, err
}

// OrderAddProductDomainTransfer post a new domain transfer in your cart
func (c *Client) OrderAddProductDomainTransfer(cartID string, orderPostDomainReq OrderPostDomainReq) (*OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	domainTransferItem := &OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domainTransfer", cartID), orderPostDomainReq, domainTransferItem)
	return domainTransferItem, err
}

// OrderGetProductDomainTransferOptions get informations about domain name transfer options
func (c *Client) OrderGetProductDomainTransferOptions(cartID string, domain string) ([]OrderCartGenericOptionDefinition, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	options := []OrderCartGenericOptionDefinition{}
	err := c.Get(queryEscape("/order/cart/%s/domainTransfer/options?domain=%s", cartID, domain), options)

	return options, err
}

// OrderAddProductDomainTransferOption post an option on a domain transfer item
func (c *Client) OrderAddProductDomainTransferOption(cartID string, orderPostDomainReq OrderPostDomainReq) (*OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	optionItem := &OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domainTransfer/options", cartID), orderPostDomainReq, optionItem)
	return optionItem, err
}

// OrderGetProductDomainRestore get products for a domain restore
func (c *Client) OrderGetProductDomainRestore(cartID string, domain string) ([]OrderCartGenericProductDefinition, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	domainRestoreProducts := []OrderCartGenericProductDefinition{}
	err := c.Get(queryEscape("/order/cart/%s/domainRestore?domain=%s", cartID, domain), domainRestoreProducts)
	return domainRestoreProducts, err
}

// OrderGetProductDomainPacks get informations about domain packs
func (c *Client) OrderGetProductDomainPacks(cartID string, domain string) ([]OrderCartDomainPacksProductInformation, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	domainPacksProducts := []OrderCartDomainPacksProductInformation{}
	err := c.Get(queryEscape("/order/cart/%s/domainPacks?domain=%s", cartID, domain), domainPacksProducts)
	return domainPacksProducts, err
}

// OrderPostProductDomainPacks post a new domain packs in your cart
func (c *Client) OrderPostProductDomainPacks(cartID string, orderPostDomainPacksReq OrderPostDomainPacksReq) (*OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("Error 404: \"Invalid Cart ID\"")
	}
	domainPacksItem := &OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domainTransfer", cartID), orderPostDomainPacksReq, domainPacksItem)
	return domainPacksItem, err
}
