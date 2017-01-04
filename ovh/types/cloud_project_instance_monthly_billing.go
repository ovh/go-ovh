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

// Instance monthly billing details
type CloudProjectInstanceMonthlyBilling struct {

	// Monthly billing activation date
	ActivatedOn time.Time `json:"activatedOn,omitempty"`

	Cost OrderPrice `json:"cost,omitempty"`
}
