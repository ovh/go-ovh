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

// Instance
type CloudInstanceInstance struct {

	// Instance creation date
	Created time.Time `json:"created,omitempty"`

	// Instance flavor id
	FlavorId string `json:"flavorId,omitempty"`

	// Instance id
	Id string `json:"id,omitempty"`

	// Instance image id
	ImageId string `json:"imageId,omitempty"`

	// Instance IP addresses
	IpAddresses []CloudInstanceIpAddress `json:"ipAddresses,omitempty"`

	MonthlyBilling CloudInstanceMonthlyBilling `json:"monthlyBilling,omitempty"`

	// Instance name
	Name string `json:"name,omitempty"`

	// Instance id
	Region string `json:"region,omitempty"`

	// Instance ssh key id
	SshKeyId string `json:"sshKeyId,omitempty"`

	// Instance status
	Status string `json:"status,omitempty"`
}
