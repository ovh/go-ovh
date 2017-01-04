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

// Credit details
type CloudProjectNewProjectCredit struct {

	// Credit description
	Description string `json:"description,omitempty"`

	// Credit id
	Id int64 `json:"id,omitempty"`

	// Use credits on following products
	Products []string `json:"products,omitempty"`

	TotalCredit OrderPrice `json:"total_credit,omitempty"`

	Validity CloudCommonVoucherValidity `json:"validity,omitempty"`
}
