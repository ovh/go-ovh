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

// VrackDedicatedServer vrack dedicated server interfaces
type VrackDedicatedServer struct {

	// DedicatedServer Dedicated Server
	DedicatedServer string `json:"dedicatedServer,omitempty"`

	// Vrack vrack name
	Vrack string `json:"vrack,omitempty"`
}