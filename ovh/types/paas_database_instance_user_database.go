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

// Databases linked to an user
type PaasDatabaseInstanceUserDatabase struct {

	// Database's name linked to this user
	DatabaseName string `json:"databaseName,omitempty"`

	// The grantId associated to this userName for this databaseName
	GrantId string `json:"grantId,omitempty"`

	// Grant of this user for this database
	GrantType string `json:"grantType,omitempty"`
}
