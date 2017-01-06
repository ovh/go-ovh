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

// SSLGatewayDomain Domain attached to an SSL Gateway
type SSLGatewayDomain struct {

	// Domain Domain name attached to your SSL Gateway
	Domain string `json:"domain,omitempty"`

	// ID Id of your domain
	ID int64 `json:"id,omitempty"`

	// State Domain state
	State string `json:"state,omitempty"`
}