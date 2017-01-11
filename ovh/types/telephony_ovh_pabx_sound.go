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

// TelephonyOvhPabxSound The PABX sounds
type TelephonyOvhPabxSound struct {

	// Name The sound filename
	Name string `json:"name,omitempty"`

	SoundID int64 `json:"soundId,omitempty"`
}
