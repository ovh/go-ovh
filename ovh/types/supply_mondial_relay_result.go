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

// Status and Mondial Relay Point Details
type SupplyMondialRelayResult struct {

	// Reference address for finding RelayPoints
	ReferenceAddress string `json:"referenceAddress,omitempty"`

	// Array of relay points
	RelayPoints []SupplyMondialRelay `json:"relayPoints,omitempty"`
}
