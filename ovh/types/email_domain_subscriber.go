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

// EmailDomainSubscriber Subscribers List
type EmailDomainSubscriber struct {
	Domain string `json:"domain,omitempty"`

	Email string `json:"email,omitempty"`

	Mailinglist string `json:"mailinglist,omitempty"`
}