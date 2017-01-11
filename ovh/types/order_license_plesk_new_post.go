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

// OrderLicensePleskNewPost ...
type OrderLicensePleskNewPost struct {
	Antivirus string `json:"antivirus,omitempty"`

	ApplicationSet string `json:"applicationSet,omitempty"`

	DomainNumber string `json:"domainNumber,omitempty"`

	IP string `json:"ip,omitempty"`

	LanguagePackNumber string `json:"languagePackNumber,omitempty"`

	Powerpack bool `json:"powerpack,omitempty"`

	ResellerManagement bool `json:"resellerManagement,omitempty"`

	ServiceType string `json:"serviceType,omitempty"`

	Version string `json:"version,omitempty"`

	WordpressToolkit bool `json:"wordpressToolkit,omitempty"`
}
