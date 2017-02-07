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

// DockerSLAveFlavor Attributes of the slave flavor
type DockerSLAveFlavor struct {

	// Bandwidth The network bandwidth, in Mbps
	Bandwidth int64 `json:"bandwidth,omitempty"`

	// CPUs The amount of (v)CPUs
	CPUs int64 `json:"cpus,omitempty"`

	// Disk The disk size, in GB
	Disk int64 `json:"disk,omitempty"`

	// DiskHa Wether the disk is HA (stored in Ceph) or local (SSD)
	DiskHa bool `json:"diskHa,omitempty"`

	// ID The flavor UUID
	ID string `json:"id,omitempty"`

	// IsVM Whether the flavor is an Openstack or dedicated flavor
	IsVM bool `json:"isVm,omitempty"`

	// RAM The amount of RAM, in MB
	RAM int64 `json:"ram,omitempty"`
}
