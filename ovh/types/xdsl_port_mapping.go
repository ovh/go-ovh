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

// XdslPortMapping Port Mappings
type XdslPortMapping struct {

	// AllowedRemoteIP An ip which will access to the defined rule. Default : no restriction applied
	AllowedRemoteIP string `json:"allowedRemoteIp,omitempty"`

	// Description Description of the Port Mapping
	Description string `json:"description,omitempty"`

	// ExternalPortEnd The last port of the interval on the External Client that will get the connections
	ExternalPortEnd int64 `json:"externalPortEnd,omitempty"`

	// ExternalPortStart External Port that the modem will listen on. List of externalPorts not available for now in the API : 8, 21, 68, 5060, 21800-21805, 51005
	ExternalPortStart int64 `json:"externalPortStart,omitempty"`

	// ID ID of the port mapping entry
	ID int64 `json:"id,omitempty"`

	// InternalClient The IP address of the destination of the packets
	InternalClient string `json:"internalClient,omitempty"`

	// InternalPort The port on the Internal Client that will get the connections
	InternalPort int64 `json:"internalPort,omitempty"`

	// Name Name of the port mapping entry
	Name string `json:"name,omitempty"`

	// Protocol Protocol of the port mapping (TCP / UDP)
	Protocol string `json:"protocol,omitempty"`

	// TaskID ID of the ongoing todo (NULL if none)
	TaskID int64 `json:"taskId,omitempty"`
}
