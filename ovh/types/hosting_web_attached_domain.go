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

// Virtual service
type HostingWebAttachedDomain struct {

	// Is linked to the hosting cdn
	Cdn string `json:"cdn,omitempty"`

	// Domain linked (fqdn)
	Domain string `json:"domain,omitempty"`

	// Firewall state for this path
	Firewall string `json:"firewall,omitempty"`

	// IP location of the domain linked
	IpLocation string `json:"ipLocation,omitempty"`

	// Put domain for separate the logs
	OwnLog string `json:"ownLog,omitempty"`

	// Domain path, relative to your home directory
	Path string `json:"path,omitempty"`

	// Put domain in ssl certificate
	Ssl bool `json:"ssl,omitempty"`
}
