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

// PcaFile cloud archives files
type PcaFile struct {

	// MD5 File MD5 hash
	MD5 string `json:"MD5,omitempty"`

	// SHA1 File SHA1 hash
	SHA1 string `json:"SHA1,omitempty"`

	// SHA256 File SHA256 hash
	SHA256 string `json:"SHA256,omitempty"`

	// ID File id
	ID string `json:"id,omitempty"`

	// Name File name
	Name string `json:"name,omitempty"`

	// Size File size, in bytes
	Size int64 `json:"size,omitempty"`

	// State File state
	State string `json:"state,omitempty"`

	// TType File type
	TType string `json:"type,omitempty"`
}