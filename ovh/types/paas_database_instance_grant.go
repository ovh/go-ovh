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

// PaasDatabaseInstanceGrant Grants
type PaasDatabaseInstanceGrant struct {

	// CreationDate Creation date of the grant
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// DatabaseName Database's name for this grant
	DatabaseName string `json:"databaseName,omitempty"`

	// GrantID Grant id
	GrantID string `json:"grantId,omitempty"`

	// GrantType Grant type
	GrantType string `json:"grantType,omitempty"`

	// LastUpdate The last update date of this grant
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// Status Grant status
	Status string `json:"status,omitempty"`

	// TaskID The id of the task working on this object
	TaskID string `json:"taskId,omitempty"`

	// UserName User name to grant
	UserName string `json:"userName,omitempty"`
}
