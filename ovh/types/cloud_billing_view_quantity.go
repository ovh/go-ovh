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

// Quantity
type CloudBillingViewQuantity struct {

	// Quantity unit
	Unit string `json:"unit,omitempty"`

	// Quantity value
	Value float64 `json:"value,omitempty"`
}
