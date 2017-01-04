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

// Hosting users logs
type HostingWebUserLogs struct {

	// Date of the user creation
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Description field for you
	Description string `json:"description,omitempty"`

	// The userLogs login used to connect to logs.ovh.net
	Login string `json:"login,omitempty"`
}
