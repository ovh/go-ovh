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

// Task on a SSL
type SslOperation struct {

	// Completion date
	DoneDate time.Time `json:"doneDate,omitempty"`

	// Task function name
	Function string `json:"function,omitempty"`

	// Task last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Task Creation date
	StartDate time.Time `json:"startDate,omitempty"`

	// Task status
	Status string `json:"status,omitempty"`

	TaskId int64 `json:"taskId,omitempty"`
}
