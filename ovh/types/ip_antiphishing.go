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

// Phishing URLs hosted on your IP
type IpAntiphishing struct {

	// Date of the event
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Internal ID of the phishing entry
	Id int64 `json:"id,omitempty"`

	// IP address hosting the phishing URL
	IpOnAntiphishing string `json:"ipOnAntiphishing,omitempty"`

	// Current state of the phishing
	State string `json:"state,omitempty"`

	// Phishing URL
	UrlPhishing string `json:"urlPhishing,omitempty"`
}
