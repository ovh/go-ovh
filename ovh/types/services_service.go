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

// ServicesService Details about a Service
type ServicesService struct {

	// CanDeleteAtExpiration Indicates that the service can be set up to be deleted at expiration
	CanDeleteAtExpiration bool `json:"canDeleteAtExpiration,omitempty"`

	ContactAdmin string `json:"contactAdmin,omitempty"`

	ContactBilling string `json:"contactBilling,omitempty"`

	ContactTech string `json:"contactTech,omitempty"`

	Creation *time.Time `json:"creation,omitempty"`

	Domain string `json:"domain,omitempty"`

	EngagedUpTo *time.Time `json:"engagedUpTo,omitempty"`

	Expiration *time.Time `json:"expiration,omitempty"`

	// PossibleRenewPeriod All the possible renew period of your service in month
	PossibleRenewPeriod []int64 `json:"possibleRenewPeriod,omitempty"`

	Renew *ServiceRenewType `json:"renew,omitempty"`

	RenewalType string `json:"renewalType,omitempty"`

	Status string `json:"status,omitempty"`
}