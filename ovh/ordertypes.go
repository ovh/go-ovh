package ovh

import "time"

// OrderCart is a go representation of Cart instance
type OrderCart struct {
	CartID      string     `json:"cartId,omitempty"`
	Expire      *time.Time `json:"expire,omitempty"`
	Description string     `json:"description,omitempty"`
	ReadOnly    bool       `json:"readOnly,omitempty"`
	Items       []int      `json:"items,omitempty"`
}

// Order is a go representation of Order instance
type Order struct {
	OrderID   int             `json:"orderId,omitempty"`
	URL       string          `json:"url,omitempty"`
	Details   []OrderDetail   `json:"details,omitempty"`
	Contracts []OrderContract `json:"contracts,omitempty"`
	Prices    OrderPrices     `json:"prices,omitempty"`
}

// OrderDetail is a go representation of OrderDetail instance
type OrderDetail struct {
	Domain      string     `json:"domain,omitempty"`
	TotalPrice  OrderPrice `json:"totalPrice,omitempty"`
	DetailType  string     `json:"detailType,omitempty"`
	Quantity    int        `json:"quantity,omitempty"`
	UnitPrice   OrderPrice `json:"unitPrice,omitempty"`
	Description string     `json:"description,omitempty"`
}

// OrderContract is a go representation of OrderContract instance
type OrderContract struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	Content string `json:"content,omitempty"`
}

// OrderPrices is a go representation of OrderPrices instance
type OrderPrices struct {
	WithoutTax OrderPrice `json:"withoutTax,omitempty"`
	Tax        OrderPrice `json:"tax,omitempty"`
	WithTax    OrderPrice `json:"withTax,omitempty"`
}

// OrderPrice is a go representation of OrderPrice instance
type OrderPrice struct {
	CurrencyCode string  `json:"currencyCode,omitempty"`
	Value        float32 `json:"value,omitempty"`
	Text         string  `json:"text,omitempty"`
}

//OrderCartItem is a representation of a cart item
type OrderCartItem struct {
	OfferID        string                  `json:"offerId,omitempty"`
	Options        []int                   `json:"options,omitempty"`
	ParentItemID   int                     `json:"parentItemId,omitempty"`
	ProductID      string                  `json:"productId,omitempty"`
	Duration       string                  `json:"duration,omitempty"`
	Settings       OrderCartDomainSettings `json:"settings,omitempty"`
	ItemID         int                     `json:"itemID,omitempty"`
	Configurations []int                   `json:"configurations,omitempty"`
	CartID         string                  `json:"cartId,omitempty"`
	Prices         []OrderCartPrice        `json:"prices,omitempty"`
}

//OrderCartPrice is a representation of a cart price
type OrderCartPrice struct {
	Label string     `json:"label,omitempty"`
	Price OrderPrice `json:"price,omitempty"`
}

//OrderCartDomainSettings is a representation of domain name order properties
type OrderCartDomainSettings struct {
	Domain string `json:"domain,omitempty"`
	Phase string `json:"phase,omitempty"`
	Offer string `json:"offer,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}

//OrderCartConfigurationItem is a representation of a configuration item for personalizing product
type OrderCartConfigurationItem struct {
	ID    int    `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

//OrderCartProductInformation is a representation of a product information
type OrderCartProductInformation struct {
	DeliveryTime   string                               `json:"deliveryTime"`
	ProductID      string                               `json:"productId,omitempty"`
	Offer          string                               `json:"offer,omitempty"`
	OfferID        string                               `json:"offerId,omitempty"`
	Duration       []string                             `json:"duration,omitempty"`
	Orderable      bool                                 `json:"orderable"`
	Configurations []OrderCartConfigurationRequirements `json:"configurations"`
	Phase          string                               `json:"phase,omitempty"`
	QuantityMax    int                                  `json:"quantityMax,omitempty"`
	Prices         []OrderCartPrice                     `json:"prices,omitempty"`
}

//OrderCartConfigurationRequirements is a representation of a configuration requirements
type OrderCartConfigurationRequirements struct {
	Required bool     `json:"required"`
	Fields   []string `json:"fields"`
	Label    string   `json:"label,omitempty"`
	Type     string   `json:"type,omitempty"`
}

//OrderCartGenericOptionDefinition is a representation of a generic option
type OrderCartGenericOptionDefinition struct {
	ProductType string                           `json:"productType,omitempty"`
	ProductName string                           `json:"productName,omitempty"`
	PlanCode    string                           `json:"planCode,omitempty"`
	Mandatory   bool                             `json:"mandatory,omitempty"`
	Prices      []OrderCartGenericProductPricing `json:"prices,omitempty"`
	Family      string                           `json:"family,omitempty"`
	Exclusive   bool                             `json:"exclusive,omitempty"`
}

//OrderCartGenericProductDefinition is a representation of a generic product
type OrderCartGenericProductDefinition struct {
	ProductName string                           `json:"productName,omitempty"`
	ProductType string                           `json:"productType,omitempty"`
	PlanCode    string                           `json:"planCode,omitempty"`
	Prices      []OrderCartGenericProductPricing `json:"prices,omitempty"`
}

//OrderCartGenericProductPricing is a representation of a generic product pricing
type OrderCartGenericProductPricing struct {
	PriceInUcents   int        `json:"priceInUcents,omitempty"`
	Capacities      string     `json:"capacities,omitempty"`
	PricingMode     string     `json:"pricingMode,omitempty"`
	Duration        string     `json:"duration,omitempty"`
	Description     string     `json:"description,omitempty"`
	Interval        int        `json:"interval,omitempty"`
	MinimumRepeat   int        `json:"minimumRepeat,omitempty"`
	PricingType     string     `json:"pricingType,omitempty"`
	MaximumQuantity int        `json:"maximumQuantity,omitempty"`
	MaximumRepeat   int        `json:"maximumRepeat,omitempty"`
	MinimumQuantity int        `json:"minimumQuantity,omitempty"`
	Price           OrderPrice `json:"price,omitempty"`
}

//OrderCartDomainPacksProductInformation is a representation of a domain packs product
type OrderCartDomainPacksProductInformation struct {
	PlanCode    string                           `json:"planCode,omitempty"`
	Description string                           `json:"description,omitempty"`
	Prices      []OrderCartGenericProductPricing `json:"prices,omitempty"`
}
