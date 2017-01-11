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

// CloudStorageContainer Container
type CloudStorageContainer struct {

	// ID Storage id
	ID string `json:"id,omitempty"`

	// Name Storage name
	Name string `json:"name,omitempty"`

	Region string `json:"region,omitempty"`

	// StoredBytes Total bytes stored
	StoredBytes int64 `json:"storedBytes,omitempty"`

	// StoredObjects Total objects stored
	StoredObjects int64 `json:"storedObjects,omitempty"`
}
