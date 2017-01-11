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

// DeskaasTask Operation on a Desktop As A Service component
type DeskaasTask struct {

	// Description Current progress description
	Description string `json:"description,omitempty"`

	// LastModificationDate Task last modification date
	LastModificationDate *time.Time `json:"lastModificationDate,omitempty"`

	// Name Task name
	Name string `json:"name,omitempty"`

	// Progress Current progress
	Progress int64 `json:"progress,omitempty"`

	// State Current Task state
	State string `json:"state,omitempty"`

	// TaskID Task id
	TaskID int64 `json:"taskId,omitempty"`
}
