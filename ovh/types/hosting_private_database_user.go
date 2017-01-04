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

import (
	"time"
)

// Users
type HostingPrivateDatabaseUser struct {

	// Creation date of the user
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Databases granted for this user
	Databases []HostingPrivateDatabaseUserDatabase `json:"databases,omitempty"`

	// User name used to connect to your databases
	UserName string `json:"userName,omitempty"`
}
