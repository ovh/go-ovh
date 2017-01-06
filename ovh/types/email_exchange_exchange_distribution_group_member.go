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

// EmailExchangeDistributionGroupMember Mailing list members
type EmailExchangeDistributionGroupMember struct {

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// MemberAccountID Member account id
	MemberAccountID int64 `json:"memberAccountId,omitempty"`

	// MemberContactID Member account id
	MemberContactID int64 `json:"memberContactId,omitempty"`

	// MemberEmailAddress Member account primaryEmailAddress
	MemberEmailAddress string `json:"memberEmailAddress,omitempty"`

	// TaskPendingID Pending task id
	TaskPendingID int64 `json:"taskPendingId,omitempty"`
}