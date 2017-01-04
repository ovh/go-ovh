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

// Configuration of the firewall
type MonitoringFirewallConfig struct {

	// List of allowed networks to the LiveStatus API
	Livestatus []string `json:"livestatus,omitempty"`

	// List of allowed networks to the NSCA receiver
	Nsca []string `json:"nsca,omitempty"`

	// List of allowed networks to the Thruk web interface
	Thruk []string `json:"thruk,omitempty"`
}
