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

// Spam and Antyvirus configuration
type EmailExchangeSpamAndVirusConfiguration struct {

	// Check DKIM of message
	CheckDKIM bool `json:"checkDKIM,omitempty"`

	// Check SPF of message
	CheckSPF bool `json:"checkSPF,omitempty"`

	// If message is a spam delete it
	DeleteSpam bool `json:"deleteSpam,omitempty"`

	// If message is a virus delete it
	DeleteVirus bool `json:"deleteVirus,omitempty"`

	// If message is a spam or virus put in junk. Overridden by deleteSpam or deleteVirus
	PutInJunk bool `json:"putInJunk,omitempty"`

	// If message is a spam change its subject
	TagSpam bool `json:"tagSpam,omitempty"`

	// If message is a virus change its subject
	TagVirus bool `json:"tagVirus,omitempty"`
}
