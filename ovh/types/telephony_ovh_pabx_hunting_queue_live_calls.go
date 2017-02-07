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

// TelephonyOvhPabxHuntingQueueLiveCalls Live statistics of the queue
type TelephonyOvhPabxHuntingQueueLiveCalls struct {

	// Agent Name or number of the agent who answered the call
	Agent string `json:"agent,omitempty"`

	// Answered Answer date of the call
	Answered *time.Time `json:"answered,omitempty"`

	// Begin Begin date of the call
	Begin *time.Time `json:"begin,omitempty"`

	// CallerIDName Name of the caller (or anonymous if unknown)
	CallerIDName string `json:"callerIdName,omitempty"`

	// CallerIDNumber Phone number of the caller (or anonymous if unknown)
	CallerIDNumber string `json:"callerIdNumber,omitempty"`

	// DestinationNumber Phone number called (in case of an outgoing call)
	DestinationNumber string `json:"destinationNumber,omitempty"`

	// End End date of the call
	End *time.Time `json:"end,omitempty"`

	ID int64 `json:"id,omitempty"`

	// OnHold Whether or not the call is on hold
	OnHold bool `json:"onHold,omitempty"`

	// Queue Name of the queue of the call
	Queue string `json:"queue,omitempty"`

	// State Current state of the call
	State string `json:"state,omitempty"`
}
