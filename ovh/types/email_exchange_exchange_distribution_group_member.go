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

// Mailing list members
type EmailExchangeExchangeDistributionGroupMember struct {

	// Creation date
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Member account id
	MemberAccountId int64 `json:"memberAccountId,omitempty"`

	// Member account id
	MemberContactId int64 `json:"memberContactId,omitempty"`

	// Member account primaryEmailAddress
	MemberEmailAddress string `json:"memberEmailAddress,omitempty"`

	// Pending task id
	TaskPendingId int64 `json:"taskPendingId,omitempty"`
}
