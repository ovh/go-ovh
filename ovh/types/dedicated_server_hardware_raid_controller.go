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

// RAID controller configuration
type DedicatedServerHardwareRaidController struct {

	// Connected disk type
	Disks []DedicatedServerHardwareRaidDiskGroup `json:"disks,omitempty"`

	// Raid controler model
	Model string `json:"model,omitempty"`

	// Raid controler type
	Type_ string `json:"type,omitempty"`
}
