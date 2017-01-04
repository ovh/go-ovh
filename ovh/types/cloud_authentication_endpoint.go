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

// Endpoint
type CloudAuthenticationEndpoint struct {

	Id string `json:"id,omitempty"`

	Interface_ string `json:"interface,omitempty"`

	LegacyEndpointId string `json:"legacy_endpoint_id,omitempty"`

	RegionId string `json:"region_id,omitempty"`

	ServiceId string `json:"service_id,omitempty"`

	Url string `json:"url,omitempty"`
}
