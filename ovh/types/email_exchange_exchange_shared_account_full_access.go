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

// Users having full access on this shared mailbox
type EmailExchangeExchangeSharedAccountFullAccess struct {

	// Account id to give full access
	AllowedAccountId int64 `json:"allowedAccountId,omitempty"`

	// Creation date
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Pending task id
	TaskPendingId int64 `json:"taskPendingId,omitempty"`
}
