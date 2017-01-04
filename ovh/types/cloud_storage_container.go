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

// Container
type CloudStorageContainer struct {

	// Storage id
	Id string `json:"id,omitempty"`

	// Storage name
	Name string `json:"name,omitempty"`

	Region string `json:"region,omitempty"`

	// Total bytes stored
	StoredBytes int64 `json:"storedBytes,omitempty"`

	// Total objects stored
	StoredObjects int64 `json:"storedObjects,omitempty"`
}
