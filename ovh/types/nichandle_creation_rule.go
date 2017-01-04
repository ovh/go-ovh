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

// Describe all rules for a given field
type NichandleCreationRule struct {

	Mandatory bool `json:"mandatory,omitempty"`

	RegularExpression string `json:"regularExpression,omitempty"`
}
