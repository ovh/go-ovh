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

// HostingWebEmail Hosting automated emails
type HostingWebEmail struct {

	// Bounce Bounce
	Bounce int64 `json:"bounce,omitempty"`

	// Email Email used to receive errors
	Email string `json:"email,omitempty"`

	// MaxPerDay Max email to sent per day
	MaxPerDay int64 `json:"maxPerDay,omitempty"`

	// Sent Email sent since hosting creation
	Sent int64 `json:"sent,omitempty"`

	// SentToday Email sent today
	SentToday int64 `json:"sentToday,omitempty"`

	// State Email state
	State string `json:"state,omitempty"`
}
