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

// XdslLineDiagnosticPossibleValue possible value for specific answer
type XdslLineDiagnosticPossibleValue struct {

	// ID answer id
	ID int64 `json:"id,omitempty"`

	// Label answer choice string
	Label string `json:"label,omitempty"`

	// Value answer choice value
	Value string `json:"value,omitempty"`
}
