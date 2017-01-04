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

// Private database
type HostingPrivateDatabaseService struct {

	// Number of cpu on your private database
	Cpu int64 `json:"cpu,omitempty"`

	// Datacenter where this private database is located
	Datacenter string `json:"datacenter,omitempty"`

	// Set the name displayed in customer panel for your private database (max 50 chars)
	DisplayName string `json:"displayName,omitempty"`

	GraphEndpoint HostingPrivateDatabaseGraphEndpoint `json:"graphEndpoint,omitempty"`

	// URL for the graphical user interface
	GuiURL string `json:"guiURL,omitempty"`

	// Private database hostname
	Hostname string `json:"hostname,omitempty"`

	// Private database ftp hostname
	HostnameFtp string `json:"hostnameFtp,omitempty"`

	// Infrastructure where service was stored
	Infrastructure string `json:"infrastructure,omitempty"`

	// Private database ip
	Ip string `json:"ip,omitempty"`

	// Date of the last data synchronization
	LastCheck time.Time `json:"lastCheck,omitempty"`

	// Private database service port
	Port int64 `json:"port,omitempty"`

	// Private database ftp port
	PortFtp int64 `json:"portFtp,omitempty"`

	QuotaSize HostingPrivateDatabaseServiceQuotaSize `json:"quotaSize,omitempty"`

	QuotaUsed HostingPrivateDatabaseServiceQuotaUsed `json:"quotaUsed,omitempty"`

	Ram HostingPrivateDatabaseServiceRam `json:"ram,omitempty"`

	// Private database server name
	Server string `json:"server,omitempty"`

	// Service name
	ServiceName string `json:"serviceName,omitempty"`

	// Private database state
	State string `json:"state,omitempty"`

	// Private database type
	Type_ string `json:"type,omitempty"`

	// Private database version
	Version string `json:"version,omitempty"`
}
