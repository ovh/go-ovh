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

// Physical or Virtual Node
type ClusterHadoopNode struct {

	// Name of the billing profile attached to the node
	BillingProfileName string `json:"billingProfileName,omitempty"`

	// Hostname of the node
	Hostname string `json:"hostname,omitempty"`

	// IP of the Node
	Ip string `json:"ip,omitempty"`

	// Wether or not the Node is removable
	IsRemovable bool `json:"isRemovable,omitempty"`

	// Profile of the Node
	SoftwareProfile string `json:"softwareProfile,omitempty"`

	// State of the Node
	State string `json:"state,omitempty"`
}
