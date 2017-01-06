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

// LicenseVirtuozzo Your Virtuozzo license
type LicenseVirtuozzo struct {

	// ContainerNumber The amount of containers this license can manage
	ContainerNumber string `json:"containerNumber,omitempty"`

	// Creation This license creation date
	Creation *time.Time `json:"creation,omitempty"`

	// DeleteAtExpiration Shall we delete this on expiration ?
	DeleteAtExpiration bool `json:"deleteAtExpiration,omitempty"`

	// Domain The internal name of your license
	Domain string `json:"domain,omitempty"`

	// InformationKey This license Information key
	InformationKey string `json:"informationKey,omitempty"`

	// IP The ip on which this license is attached
	IP string `json:"ip,omitempty"`

	// LicenseID The license id on license provider side
	LicenseID string `json:"licenseId,omitempty"`

	// ProductKey This license product key
	ProductKey string `json:"productKey,omitempty"`

	// Status This license state
	Status string `json:"status,omitempty"`

	// Version This license version
	Version string `json:"version,omitempty"`
}