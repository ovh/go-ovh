/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

// OrderDetail Detail of an order
type OrderDetail struct {
	Description string `json:"description,omitempty"`

	DetailType string `json:"detailType,omitempty"`

	Domain string `json:"domain,omitempty"`

	Quantity int64 `json:"quantity,omitempty"`

	TotalPrice *OrderPrice `json:"totalPrice,omitempty"`

	UnitPrice *OrderPrice `json:"unitPrice,omitempty"`
}
