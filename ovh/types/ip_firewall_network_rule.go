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

// Rule on ip
type IpFirewallNetworkRule struct {

	// Action on this rule
	Action string `json:"action,omitempty"`

	CreationDate time.Time `json:"creationDate,omitempty"`

	// Destination ip for your rule
	Destination string `json:"destination,omitempty"`

	// Destination port range for your rule. Only with TCP/UDP protocol
	DestinationPort string `json:"destinationPort,omitempty"`

	// Fragments option
	Fragments bool `json:"fragments,omitempty"`

	// Network protocol
	Protocol string `json:"protocol,omitempty"`

	Rule string `json:"rule,omitempty"`

	Sequence int64 `json:"sequence,omitempty"`

	// Source ip for your rule
	Source string `json:"source,omitempty"`

	// Source port range for your rule. Only with TCP/UDP protocol
	SourcePort string `json:"sourcePort,omitempty"`

	// Current state of your rule
	State string `json:"state,omitempty"`

	// TCP option on your rule
	TcpOption string `json:"tcpOption,omitempty"`
}
