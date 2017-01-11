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

// DedicatedServerHardwareRaidDiskGroup Disk attached to a RAID controller
type DedicatedServerHardwareRaidDiskGroup struct {
	Capacity *DedicatedServerHardwareRaidConfigurationDiskSize `json:"capacity,omitempty"`

	// Names Disk names
	Names []string `json:"names,omitempty"`

	Speed *DedicatedServerHardwareRaidDiskGroupSpeed `json:"speed,omitempty"`

	// TType Disk type
	TType string `json:"type,omitempty"`
}
