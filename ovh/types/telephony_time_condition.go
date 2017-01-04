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

// Time conditions
type TelephonyTimeCondition struct {

	Day string `json:"day,omitempty"`

	HourBegin string `json:"hourBegin,omitempty"`

	HourEnd string `json:"hourEnd,omitempty"`

	Id int64 `json:"id,omitempty"`

	Policy string `json:"policy,omitempty"`

	Status string `json:"status,omitempty"`
}
