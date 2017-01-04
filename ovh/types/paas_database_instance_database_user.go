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

// User granted to a database
type PaasDatabaseInstanceDatabaseUser struct {

	// The grantId associated for this databaseName and this userName
	GrantId string `json:"grantId,omitempty"`

	// User's rights on this database
	GrantType string `json:"grantType,omitempty"`

	// User's name granted on this database
	UserName string `json:"userName,omitempty"`
}
