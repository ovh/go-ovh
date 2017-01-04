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

// Live statistics of the queue
type TelephonyOvhPabxHuntingQueueLiveCalls struct {

	// Name or number of the agent who answered the call
	Agent string `json:"agent,omitempty"`

	// Answer date of the call
	Answered time.Time `json:"answered,omitempty"`

	// Begin date of the call
	Begin time.Time `json:"begin,omitempty"`

	// Name of the caller (or anonymous if unknown)
	CallerIdName string `json:"callerIdName,omitempty"`

	// Phone number of the caller (or anonymous if unknown)
	CallerIdNumber string `json:"callerIdNumber,omitempty"`

	// Phone number called (in case of an outgoing call)
	DestinationNumber string `json:"destinationNumber,omitempty"`

	// End date of the call
	End time.Time `json:"end,omitempty"`

	Id int64 `json:"id,omitempty"`

	// Whether or not the call is on hold
	OnHold bool `json:"onHold,omitempty"`

	// Name of the queue of the call
	Queue string `json:"queue,omitempty"`

	// Current state of the call
	State string `json:"state,omitempty"`
}
