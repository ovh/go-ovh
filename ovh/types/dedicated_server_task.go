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

// DedicatedServerTask Server tasks
type DedicatedServerTask struct {

	// Comment Details of this task
	Comment string `json:"comment,omitempty"`

	// DoneDate Completion date
	DoneDate *time.Time `json:"doneDate,omitempty"`

	// Function Function name
	Function string `json:"function,omitempty"`

	// LastUpdate last update
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// StartDate Task Creation date
	StartDate *time.Time `json:"startDate,omitempty"`

	// Status Task status
	Status string `json:"status,omitempty"`

	// TaskID the id of the task
	TaskID int64 `json:"taskId,omitempty"`
}
