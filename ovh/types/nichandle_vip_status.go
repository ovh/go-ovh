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

// VIP Status by Universe
type NichandleVipStatus struct {

	// Is account VIP for Cloud Universe
	Cloud bool `json:"cloud,omitempty"`

	// Is account VIP for Dedicated Universe
	Dedicated bool `json:"dedicated,omitempty"`

	// Is account VIP for Telecom Universe
	Telecom bool `json:"telecom,omitempty"`

	// Is account VIP for Web Universe
	Web bool `json:"web,omitempty"`
}
