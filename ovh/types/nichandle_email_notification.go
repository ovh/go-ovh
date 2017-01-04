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

// Email notification
type NichandleEmailNotification struct {

	// This email body
	Body string `json:"body,omitempty"`

	// This email date
	Date time.Time `json:"date,omitempty"`

	// This email Id
	Id int64 `json:"id,omitempty"`

	// This email subject
	Subject string `json:"subject,omitempty"`
}
