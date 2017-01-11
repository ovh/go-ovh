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

// HostingResellerProductType Plesk instance type details
type HostingResellerProductType struct {

	// Consumers Number of allowed customers
	Consumers int64 `json:"consumers,omitempty"`

	// CPU Instance's cpu
	CPU string `json:"cpu,omitempty"`

	// Databases Number of allowed databases
	Databases int64 `json:"databases,omitempty"`

	// DiskSize Disk size of the instance (in GB)
	DiskSize int64 `json:"diskSize,omitempty"`

	// EmailAccounts Number of allowed email accounts
	EmailAccounts int64 `json:"emailAccounts,omitempty"`

	// RAM Instance's ram (in GB)
	RAM int64 `json:"ram,omitempty"`

	// TType Type name
	TType string `json:"type,omitempty"`

	// VCores Number of vCore
	VCores int64 `json:"vCores,omitempty"`

	// Websites Number of allowed websites
	Websites int64 `json:"websites,omitempty"`
}
