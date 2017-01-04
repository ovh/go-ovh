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

// Log entry of an auth attempt to the radius server
type XdslRadiusConnectionLog struct {

	Date time.Time `json:"date,omitempty"`

	Login string `json:"login,omitempty"`

	Message string `json:"message,omitempty"`

	State string `json:"state,omitempty"`
}
