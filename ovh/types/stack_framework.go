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

// A framework installed for a docker PaaS stack
type StackFramework struct {

	// The framework name
	AccessUrl string `json:"accessUrl,omitempty"`

	// Date of the resource creation
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// The framework UUID
	Id string `json:"id,omitempty"`

	// The framework name
	Name string `json:"name,omitempty"`

	// Date of the resource last update
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
