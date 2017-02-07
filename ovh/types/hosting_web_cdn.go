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

// HostingWebCDN CDN service
type HostingWebCDN struct {

	// Domain Domain of this CDN
	Domain string `json:"domain,omitempty"`

	// Free Option CDN free with the hosting ?
	Free bool `json:"free,omitempty"`

	// Status Status of the CDN option
	Status string `json:"status,omitempty"`

	// TType Type of the CDN
	TType string `json:"type,omitempty"`

	// Version Version of the CDN
	Version string `json:"version,omitempty"`
}
