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

// A structure describing informations about available nameserver for secondary dns 
type SecondaryDnsSecondaryDnsNameServer struct {

	// the name server
	Hostname string `json:"hostname,omitempty"`

	Ip string `json:"ip,omitempty"`

	Ipv6 string `json:"ipv6,omitempty"`
}
