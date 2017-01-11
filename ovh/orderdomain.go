package ovh

import (
	"errors"

	"github.com/runabove/go-sdk/ovh/types"
)

// OrderGetProductsDomain get products about a domain name
func (c *Client) OrderGetProductsDomain(cartID string, domain string) ([]types.OrderCartProductInformation, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	products := []types.OrderCartProductInformation{}
	err := c.Get(queryEscape("/order/cart/%s/domain?domain=%s", cartID, domain), &products)
	return products, err
}

// OrderAddProductDomain post a new domain in your cart
func (c *Client) OrderAddProductDomain(cartID string, orderCartDomainPost types.OrderCartDomainPost) (*types.OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	domainItem := &types.OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domain", cartID), orderCartDomainPost, domainItem)
	return domainItem, err
}

// OrderGetProductDomainOptions get informations about a domain name options
func (c *Client) OrderGetProductDomainOptions(cartID string, domain string) ([]types.OrderCartGenericOptionDefinition, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}

	options := []types.OrderCartGenericOptionDefinition{}
	err := c.Get(queryEscape("/order/cart/%s/domain/options?domain=%s", cartID, domain), &options)

	return options, err
}

// OrderAddProductDomainOption post an option on a domain item
func (c *Client) OrderAddProductDomainOption(cartID string, orderPostDomainOptionReq types.OrderCartDomainOptionsPost) (*types.OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	optionItem := &types.OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domain/options", cartID), orderPostDomainOptionReq, optionItem)
	return optionItem, err
}

// OrderGetProductDomainTransfer get informations about a domain transfer
func (c *Client) OrderGetProductDomainTransfer(cartID string, domain string) ([]types.OrderCartProductInformation, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}

	products := []types.OrderCartProductInformation{}
	err := c.Get(queryEscape("/order/cart/%s/domainTransfer?domain=%s", cartID, domain), &products)

	return products, err
}

// OrderAddProductDomainTransfer post a new domain transfer in your cart
func (c *Client) OrderAddProductDomainTransfer(cartID string, orderCartDomainTransferPost types.OrderCartDomainTransferPost) (*types.OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	domainTransferItem := &types.OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domainTransfer", cartID), orderCartDomainTransferPost, domainTransferItem)
	return domainTransferItem, err
}

// OrderGetProductDomainTransferOptions get informations about domain name transfer options
func (c *Client) OrderGetProductDomainTransferOptions(cartID string, domain string) ([]types.OrderCartGenericOptionDefinition, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	options := []types.OrderCartGenericOptionDefinition{}
	err := c.Get(queryEscape("/order/cart/%s/domainTransfer/options?domain=%s", cartID, domain), options)

	return options, err
}

// OrderAddProductDomainTransferOption post an option on a domain transfer item
func (c *Client) OrderAddProductDomainTransferOption(cartID string, orderCartDomainTransferOptionsPost types.OrderCartDomainTransferOptionsPost) (*types.OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	optionItem := &types.OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domainTransfer/options", cartID), orderCartDomainTransferOptionsPost, optionItem)
	return optionItem, err
}

// OrderGetProductDomainRestore get products for a domain restore
func (c *Client) OrderGetProductDomainRestore(cartID string, domain string) ([]types.OrderCartGenericProductDefinition, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}

	domainRestoreProducts := []types.OrderCartGenericProductDefinition{}
	err := c.Get(queryEscape("/order/cart/%s/domainRestore?domain=%s", cartID, domain), &domainRestoreProducts)

	return domainRestoreProducts, err
}

// OrderGetProductDomainPacks get informations about domain packs
func (c *Client) OrderGetProductDomainPacks(cartID string, domain string) ([]types.OrderCartDomainPacksProductInformation, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}

	domainPacksProducts := []types.OrderCartDomainPacksProductInformation{}
	err := c.Get(queryEscape("/order/cart/%s/domainPacks?domain=%s", cartID, domain), &domainPacksProducts)
	return domainPacksProducts, err
}

// OrderPostProductDomainPacks post a new domain packs in your cart
func (c *Client) OrderPostProductDomainPacks(cartID string, orderPostDomainPacksReq types.OrderCartDomainPacksPost) (*types.OrderCartItem, error) {
	if cartID == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	domainPacksItem := &types.OrderCartItem{}
	err := c.Post(queryEscape("/order/cart/%s/domainTransfer", cartID), orderPostDomainPacksReq, domainPacksItem)
	return domainPacksItem, err
}

// OrderGetCartServiceOptions get service options on a domain
func (c *Client) OrderGetCartServiceOptions(domain string) ([]types.OrderCartGenericOptionDefinition, error) {
	if domain == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	cartServiceOptions := []types.OrderCartGenericOptionDefinition{}
	err := c.Get(queryEscape("/order/cartServiceOption/domain/%s", domain), &cartServiceOptions)

	return cartServiceOptions, err
}

// OrderAddCartServiceOption add a service options on your cart
func (c *Client) OrderAddCartServiceOption(domain string, orderCartServiceOptionDomainPost types.OrderCartServiceOptionDomainPost) (*types.OrderCartItem, error) {
	if domain == "" {
		return nil, errors.New("CartID parameter must not be empty")
	}
	item := &types.OrderCartItem{}
	err := c.Post(queryEscape("/order/cartServiceOption/domain/%s", domain), orderCartServiceOptionDomainPost, item)
	return item, err
}
