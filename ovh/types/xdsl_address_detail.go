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

// All components of an address
type XdslAddressDetail struct {

	Building string `json:"building,omitempty"`

	City string `json:"city,omitempty"`

	Door string `json:"door,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Floor string `json:"floor,omitempty"`

	// Identifier of the city
	InseeCode string `json:"inseeCode,omitempty"`

	LastName string `json:"lastName,omitempty"`

	NumberStreet string `json:"numberStreet,omitempty"`

	Residence string `json:"residence,omitempty"`

	// Identifier of the street
	RivoliCode string `json:"rivoliCode,omitempty"`

	Stairs string `json:"stairs,omitempty"`

	Street string `json:"street,omitempty"`

	ZipCode string `json:"zipCode,omitempty"`
}
