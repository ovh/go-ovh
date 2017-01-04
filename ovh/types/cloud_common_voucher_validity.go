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

// Voucher validity range
type CloudCommonVoucherValidity struct {

	// Valid from
	From time.Time `json:"from,omitempty"`

	// Valid to
	To time.Time `json:"to,omitempty"`
}
