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

// Windows web Hosting
type HostingWindowsService struct {

	// Hosting offer
	Offer string `json:"offer,omitempty"`

	// Hosting's OS
	OperatingSystem string `json:"operatingSystem,omitempty"`

	// Hosting resource type
	ResourceType string `json:"resourceType,omitempty"`

	// Service name
	ServiceName string `json:"serviceName,omitempty"`
}
