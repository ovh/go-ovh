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

// Usage information for current month on your project
type CloudProjectBandwidthStorageUsage struct {

	// Downloaded bytes from your containers
	DownloadedBytes int64 `json:"downloadedBytes,omitempty"`

	// Region
	Region string `json:"region,omitempty"`

	Total OrderPrice `json:"total,omitempty"`
}
