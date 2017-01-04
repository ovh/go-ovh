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

import (
	"time"
)

// InstanceDetail
type CloudInstanceInstanceDetail struct {

	// Instance creation date
	Created time.Time `json:"created,omitempty"`

	Flavor CloudFlavorFlavor `json:"flavor,omitempty"`

	// Instance id
	Id string `json:"id,omitempty"`

	Image CloudImageImage `json:"image,omitempty"`

	// Instance IP addresses
	IpAddresses []CloudInstanceIpAddress `json:"ipAddresses,omitempty"`

	MonthlyBilling CloudInstanceMonthlyBilling `json:"monthlyBilling,omitempty"`

	// Instance name
	Name string `json:"name,omitempty"`

	// Instance id
	Region string `json:"region,omitempty"`

	SshKey CloudSshkeySshKeyDetail `json:"sshKey,omitempty"`

	// Instance status
	Status string `json:"status,omitempty"`
}
