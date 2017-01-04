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

// A structure describing the Resource's new price
type DedicatedCloudResourceNewPricesEntry struct {

	BillingType string `json:"billingType,omitempty"`

	Changed bool `json:"changed,omitempty"`

	Name string `json:"name,omitempty"`

	NewPrice OrderPrice `json:"newPrice,omitempty"`

	OldPrice OrderPrice `json:"oldPrice,omitempty"`

	ResourceType string `json:"resourceType,omitempty"`
}
