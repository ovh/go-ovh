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

// Sms reach list
type SmsException struct {

	// The abreviated country code.
	CountrySuffixe string `json:"countrySuffixe,omitempty"`

	// The exception message
	Messages []string `json:"messages,omitempty"`

	// The list of operators impacted.
	Operators string `json:"operators,omitempty"`

	// The type of routing restriction imposed by the operator
	RestrictionCode string `json:"restrictionCode,omitempty"`

	// The substitution sender used to bypass operator filter
	Substitution string `json:"substitution,omitempty"`
}
