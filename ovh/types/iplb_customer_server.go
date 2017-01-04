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

// Server
type IplbCustomerServer struct {

	// Address of your server
	Address string `json:"address,omitempty"`

	// Id of your server
	Id int64 `json:"id,omitempty"`

	// Status attached to your server
	Status string `json:"status,omitempty"`

	// Type of your server
	Type_ string `json:"type,omitempty"`

	// Zone of your server
	Zone string `json:"zone,omitempty"`
}
