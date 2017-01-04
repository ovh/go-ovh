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

// Get protocol status on that mailbox
type EmailExchangeExchangeAccountProtocol struct {

	// IMAP protocol enabled on that mailbox
	IMAP bool `json:"IMAP,omitempty"`

	// POP protocol enabled on that mailbox
	POP bool `json:"POP,omitempty"`

	// Mobile access enabled on that mailbox
	ActiveSync bool `json:"activeSync,omitempty"`

	// Creation date
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Last update date
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Pending task id
	TaskPendingId int64 `json:"taskPendingId,omitempty"`

	// Web mail enabled on that mailbox
	WebMail bool `json:"webMail,omitempty"`
}
