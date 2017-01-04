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

// Representation of a configuration item for personalizing product
type OrderCartConfigurationItem struct {

	// Configuration ID
	Id int64 `json:"id,omitempty"`

	// Identifier of the resource
	Label string `json:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	Value string `json:"value,omitempty"`
}
