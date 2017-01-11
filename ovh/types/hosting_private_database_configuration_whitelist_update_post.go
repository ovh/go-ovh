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

// HostingPrivateDatabaseConfigurationWhitelistUpdatePost ...
type HostingPrivateDatabaseConfigurationWhitelistUpdatePost struct {
	Comment string `json:"comment,omitempty"`

	Service bool `json:"service,omitempty"`

	Sftp bool `json:"sftp,omitempty"`
}
