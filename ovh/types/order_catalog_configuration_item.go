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

// Composition of a configuration
type OrderCatalogConfigurationItem struct {

	// Default value of the configuration if not provided
	DefaultValue string `json:"defaultValue,omitempty"`

	// Indicates if the configuration is free for writing (true) or have to follow an enum (false - have to follow values field)
	IsCustom bool `json:"isCustom,omitempty"`

	// Indicates if configuration is required
	IsMandatory bool `json:"isMandatory,omitempty"`

	// Label of the configuration
	Name string `json:"name,omitempty"`

	// Values allowed if configuration isn't custom
	Values []string `json:"values,omitempty"`
}
