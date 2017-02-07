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

// CloudVolumeSnapshot Snapshot
type CloudVolumeSnapshot struct {

	// CreationDate Snapshot creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// Description Snapshot description
	Description string `json:"description,omitempty"`

	// ID Snapshot id
	ID string `json:"id,omitempty"`

	// Name Snapshot name
	Name string `json:"name,omitempty"`

	// Region Snapshot region
	Region string `json:"region,omitempty"`

	// Size Snapshot size
	Size int64 `json:"size,omitempty"`

	// Status Snapshot status
	Status string `json:"status,omitempty"`

	// VolumeID Volume source id
	VolumeID string `json:"volumeId,omitempty"`
}
