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

// IP tasks
type IpIpTask struct {

	// Details of this task
	Comment string `json:"comment,omitempty"`

	Destination IpRoutedTo `json:"destination,omitempty"`

	// Completion date
	DoneDate time.Time `json:"doneDate,omitempty"`

	// Function name
	Function string `json:"function,omitempty"`

	// last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Task Creation date
	StartDate time.Time `json:"startDate,omitempty"`

	// Task status
	Status string `json:"status,omitempty"`

	// the id of the task
	TaskId int64 `json:"taskId,omitempty"`
}
