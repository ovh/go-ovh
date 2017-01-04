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

// Traffic on mitigation
type IpMitigationTraffic struct {

	// Bits per second
	Bps int64 `json:"bps,omitempty"`

	// Paquets per second
	Pps int64 `json:"pps,omitempty"`
}
