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

// FailoverIp
type CloudIpFailoverIp struct {

	// IP block
	Block string `json:"block,omitempty"`

	// Ip continent
	ContinentCode string `json:"continentCode,omitempty"`

	// Ip location
	Geoloc string `json:"geoloc,omitempty"`

	// Ip id
	Id string `json:"id,omitempty"`

	// Ip
	Ip string `json:"ip,omitempty"`

	// Current operation progress in percent
	Progress int64 `json:"progress,omitempty"`

	// Instance where ip is routed to
	RoutedTo string `json:"routedTo,omitempty"`

	// Ip status
	Status string `json:"status,omitempty"`

	// IP sub type
	SubType string `json:"subType,omitempty"`
}
