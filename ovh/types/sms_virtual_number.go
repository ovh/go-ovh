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

// SmsVirtualNumber Virtual numbers
type SmsVirtualNumber struct {

	// CountryCode The ISO formated country code of the number
	CountryCode string `json:"countryCode,omitempty"`

	// Number The virtual number
	Number string `json:"number,omitempty"`
}
