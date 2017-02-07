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

// TelephonyHistoryConsumption Previous billed consumptions
type TelephonyHistoryConsumption struct {
	Date *time.Time `json:"date,omitempty"`

	Price *OrderPrice `json:"price,omitempty"`

	PriceOutplan *OrderPrice `json:"priceOutplan,omitempty"`

	Status string `json:"status,omitempty"`
}
