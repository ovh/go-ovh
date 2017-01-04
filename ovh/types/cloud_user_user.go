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

// User
type CloudUserUser struct {

	// User creation date
	CreationDate time.Time `json:"creationDate,omitempty"`

	// User description
	Description string `json:"description,omitempty"`

	// User id
	Id int64 `json:"id,omitempty"`

	// User status
	Status string `json:"status,omitempty"`

	// Username
	Username string `json:"username,omitempty"`
}
