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

// HostingWebUserPost ...
type HostingWebUserPost struct {
	Home string `json:"home,omitempty"`

	IisRemoteRights string `json:"iisRemoteRights,omitempty"`

	Login string `json:"login,omitempty"`

	Password string `json:"password,omitempty"`

	SSHState string `json:"sshState,omitempty"`

	WebDavRights string `json:"webDavRights,omitempty"`
}
