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

// Detail of a fax campaign
type TelephonyFaxCampaignDetail struct {

	Failed []string `json:"failed,omitempty"`

	Success []string `json:"success,omitempty"`

	Todo []string `json:"todo,omitempty"`
}
