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

// Resource record
type ZoneResetRecord struct {

	FieldType string `json:"fieldType,omitempty"`

	// Resource record target
	Target string `json:"target,omitempty"`
}
