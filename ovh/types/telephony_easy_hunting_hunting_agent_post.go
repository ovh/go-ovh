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

// TelephonyEasyHuntingHuntingAgentPost ...
type TelephonyEasyHuntingHuntingAgentPost struct {
	Number string `json:"number,omitempty"`

	SimultaneousLines int64 `json:"simultaneousLines,omitempty"`

	Status string `json:"status,omitempty"`

	Timeout int64 `json:"timeout,omitempty"`

	WrapUpTime int64 `json:"wrapUpTime,omitempty"`
}
