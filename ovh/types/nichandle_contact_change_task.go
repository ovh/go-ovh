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

// Task running a contact change on a service
type NichandleContactChangeTask struct {

	// Account who asked the contact change
	AskingAccount string `json:"askingAccount,omitempty"`

	// Contacts to be changed
	ContactTypes []string `json:"contactTypes,omitempty"`

	// Date at which the contact change has been finished
	DateDone time.Time `json:"dateDone,omitempty"`

	// Date at which the request has been made
	DateRequest time.Time `json:"dateRequest,omitempty"`

	// Account to change contact from
	FromAccount string `json:"fromAccount,omitempty"`

	Id int64 `json:"id,omitempty"`

	// The service on which the task runs
	ServiceDomain string `json:"serviceDomain,omitempty"`

	// Current state of the request
	State string `json:"state,omitempty"`

	// Account to change contact to
	ToAccount string `json:"toAccount,omitempty"`
}
