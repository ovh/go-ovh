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

// Quota
type PaasTimeseriesQuota struct {

	// Current value
	Current int64 `json:"current,omitempty"`

	// Max allowed
	Max int64 `json:"max,omitempty"`

	Type_ string `json:"type,omitempty"`
}
