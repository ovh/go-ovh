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

// RouterNetworkPost ...
type RouterNetworkPost struct {
	Description string `json:"description,omitempty"`

	IPNet string `json:"ipNet,omitempty"`

	VlanTag int64 `json:"vlanTag,omitempty"`
}
