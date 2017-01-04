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

// Represents a city
type XdslEligibilityCity struct {

	// The identifier of the city
	InseeCode string `json:"inseeCode,omitempty"`

	// The name of the city
	Name string `json:"name,omitempty"`

	// The zip code of the city
	ZipCode string `json:"zipCode,omitempty"`
}
