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

// CloudBillingViewUsedCredits UsedCredits
type CloudBillingViewUsedCredits struct {

	// Details Details about credits that will be used
	Details []*CloudBillingViewUsedCredit `json:"details,omitempty"`

	// TotalCredit Total credit that will be used to pay the bill
	TotalCredit float64 `json:"totalCredit,omitempty"`
}
