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

// Your Cpanel license
type LicenseCpanelCpanel struct {

	// This license creation date
	Creation time.Time `json:"creation,omitempty"`

	// Shall we delete this on expiration ?
	DeleteAtExpiration bool `json:"deleteAtExpiration,omitempty"`

	// The internal name of your license
	Domain string `json:"domain,omitempty"`

	// The ip on which this license is attached
	Ip string `json:"ip,omitempty"`

	// The license id on license provider side
	LicenseId string `json:"licenseId,omitempty"`

	// This license state
	Status string `json:"status,omitempty"`

	// This license version
	Version string `json:"version,omitempty"`
}
