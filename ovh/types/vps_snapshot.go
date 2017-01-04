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

// Information about the snapshot of a VPS Virtual Machine
type VpsSnapshot struct {

	CreationDate time.Time `json:"creationDate,omitempty"`

	Description string `json:"description,omitempty"`
}
