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

// SaasCsp2OfficeTenant Office tenant
type SaasCsp2OfficeTenant struct {

	// Address Contact's address line
	Address string `json:"address,omitempty"`

	// City Contact's city
	City string `json:"city,omitempty"`

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// DisplayName Tenant's display name
	DisplayName string `json:"displayName,omitempty"`

	// FirstName Contact's first name
	FirstName string `json:"firstName,omitempty"`

	// LastName Contact's last name
	LastName string `json:"lastName,omitempty"`

	// Phone Primary phone number
	Phone string `json:"phone,omitempty"`

	// ServiceName Internal service name
	ServiceName string `json:"serviceName,omitempty"`

	// Status Tenant's status
	Status string `json:"status,omitempty"`

	// ZipCode Contact's zip code
	ZipCode string `json:"zipCode,omitempty"`
}
