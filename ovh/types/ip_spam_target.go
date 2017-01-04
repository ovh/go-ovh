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

// Spam's target information
type IpSpamTarget struct {

	// Timestamp when the email was sent
	Date int64 `json:"date,omitempty"`

	// IP address of the target
	DestinationIp string `json:"destinationIp,omitempty"`

	// The message-id of the email
	MessageId string `json:"messageId,omitempty"`

	// Spam score for the email
	Spamscore int64 `json:"spamscore,omitempty"`
}
