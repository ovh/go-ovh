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

// TelephonyVxml Vxml services
type TelephonyVxml struct {
	Description string `json:"description,omitempty"`

	Offers []string `json:"offers,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	ServiceType string `json:"serviceType,omitempty"`
}
