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

// DedicatedServerBackupFtpAcl Backup Ftp ACL for this server and Backup Ftp
type DedicatedServerBackupFtpAcl struct {

	// Cifs Wether to allow the CIFS (SMB) protocol for this ACL
	Cifs bool `json:"cifs,omitempty"`

	// Ftp Wether to allow the FTP protocol for this ACL
	Ftp bool `json:"ftp,omitempty"`

	// IPBlock The IP Block specific to this ACL
	IPBlock string `json:"ipBlock,omitempty"`

	// IsApplied Whether the rule has been applied on the Backup Ftp
	IsApplied bool `json:"isApplied,omitempty"`

	// LastUpdate Date of the last object modification
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// Nfs Wether to allow the NFS protocol for this ACL
	Nfs bool `json:"nfs,omitempty"`
}