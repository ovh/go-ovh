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

// Represents a time slot for a meeting
type XdslEligibilityMeetingSlot struct {

	// The end of the time slot
	EndDate time.Time `json:"endDate,omitempty"`

	// The beginning of the time slot
	StartDate time.Time `json:"startDate,omitempty"`

	// An opaque string that represents an intervention unit
	UiCode string `json:"uiCode,omitempty"`
}
