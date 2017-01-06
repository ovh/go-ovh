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

// DebtOperation Operation that happend on a debt
type DebtOperation struct {
	Amount *OrderPrice `json:"amount,omitempty"`

	// Date Date the operation took place on
	Date *time.Time `json:"date,omitempty"`

	OperationID int64 `json:"operationId,omitempty"`

	// Status Status of the operation
	Status string `json:"status,omitempty"`

	// TType Type of movement this operation represents
	TType string `json:"type,omitempty"`
}