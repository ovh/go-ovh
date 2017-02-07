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

// TelephonyHardwareOffer Informations related to a telephone offer
type TelephonyHardwareOffer struct {

	// Description The telephony description
	Description string `json:"description,omitempty"`

	// Name The telephony name
	Name string `json:"name,omitempty"`

	Price *OrderPrice `json:"price,omitempty"`

	// URL An URL to telephony details
	URL string `json:"url,omitempty"`
}
