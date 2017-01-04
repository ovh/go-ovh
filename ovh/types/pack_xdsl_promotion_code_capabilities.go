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

// Informations about a promotion code
type PackXdslPromotionCodeCapabilities struct {

	Amount OrderPrice `json:"amount,omitempty"`

	// True if the promotion code generation is available
	CanGenerate bool `json:"canGenerate,omitempty"`

	// Number of months of engagement
	Engagement int64 `json:"engagement,omitempty"`

	// Enum of the possible errors
	ReasonCodes []string `json:"reasonCodes,omitempty"`
}
