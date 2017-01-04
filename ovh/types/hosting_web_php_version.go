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

// State of available php versions for this account
type HostingWebPhpVersion struct {

	// Current support of this php version
	Support string `json:"support,omitempty"`

	// PHP version
	Version string `json:"version,omitempty"`
}
