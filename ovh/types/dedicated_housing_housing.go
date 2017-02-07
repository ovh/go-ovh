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

// DedicatedHousing Housing bay
type DedicatedHousing struct {

	// Datacenter Housing bay datacenter
	Datacenter string `json:"datacenter,omitempty"`

	// Name The name you give to the bay
	Name string `json:"name,omitempty"`

	// Network Housing bay network
	Network []*DedicatedHousingNetworkInfo `json:"network,omitempty"`

	Options *DedicatedHousingOptions `json:"options,omitempty"`

	// Rack The bay's description
	Rack string `json:"rack,omitempty"`

	// SecurityCode Bay Security code
	SecurityCode string `json:"securityCode,omitempty"`
}
