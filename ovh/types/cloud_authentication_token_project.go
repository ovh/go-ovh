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

// TokenProject
type CloudAuthenticationTokenProject struct {

	Domain CloudAuthenticationDomain `json:"domain,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}
