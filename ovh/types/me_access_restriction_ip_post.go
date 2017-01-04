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

type MeAccessRestrictionIpPost struct {

	Ip string `json:"ip,omitempty"`

	Rule string `json:"rule,omitempty"`

	Warning bool `json:"warning,omitempty"`
}
