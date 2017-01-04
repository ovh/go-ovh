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

type PackXdslExchangeLiteServicesPost struct {

	Antispam bool `json:"antispam,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Initials string `json:"initials,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Password string `json:"password,omitempty"`
}
