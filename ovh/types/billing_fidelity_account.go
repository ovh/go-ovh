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

// Balance of the fidelity account
type BillingFidelityAccount struct {

	AlertThreshold int64 `json:"alertThreshold,omitempty"`

	Balance int64 `json:"balance,omitempty"`

	CanBeCredited bool `json:"canBeCredited,omitempty"`

	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	OpenDate time.Time `json:"openDate,omitempty"`
}
