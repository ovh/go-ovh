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

// PcaBilling cloud archives billing
type PcaBilling struct {

	// Billed Billing action is billed.
	Billed bool `json:"billed,omitempty"`

	// Date Date and time the operation took place
	Date *time.Time `json:"date,omitempty"`

	// ID Billing id
	ID int64 `json:"id,omitempty"`

	// Quantity Quantity of bytes for operation
	Quantity int64 `json:"quantity,omitempty"`

	// Reference Billing reference name
	Reference string `json:"reference,omitempty"`

	// Total Total usage after operation
	Total int64 `json:"total,omitempty"`
}
