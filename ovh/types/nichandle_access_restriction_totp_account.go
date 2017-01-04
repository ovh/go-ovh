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

// TOTP Two-Factor Authentication
type NichandleAccessRestrictionTotpAccount struct {

	// The Id of the restriction
	Id int64 `json:"id,omitempty"`

	// Status of this account
	Status string `json:"status,omitempty"`
}
