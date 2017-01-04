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

// Hourly consumption of a host
type DedicatedCloudHostHourlyConsumption struct {

	Consumption DedicatedCloudHostHourlyConsumptionConsumption `json:"consumption,omitempty"`

	// Last update.
	LastUpdate time.Time `json:"lastUpdate,omitempty"`
}
