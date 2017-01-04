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

// Stream rule
type DBaaSLogsStreamRule struct {

	// Field name
	Field string `json:"field,omitempty"`

	// Invert condition
	IsInverted bool `json:"isInverted,omitempty"`

	// Field operator
	Operator string `json:"operator,omitempty"`

	// Stream rule UUID
	RuleId string `json:"ruleId,omitempty"`

	// Field value
	Value string `json:"value,omitempty"`
}
