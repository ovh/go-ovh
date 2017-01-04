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

// cloud storage billing
type PcsBilling struct {

	// Billing action is billed.
	Billed bool `json:"billed,omitempty"`

	// Date and time the operation took place
	Date time.Time `json:"date,omitempty"`

	// Billing id
	Id int64 `json:"id,omitempty"`

	// Quantity of bytes for operation
	Quantity int64 `json:"quantity,omitempty"`

	// Billing reference name
	Reference string `json:"reference,omitempty"`

	// Total usage after operation
	Total int64 `json:"total,omitempty"`
}
