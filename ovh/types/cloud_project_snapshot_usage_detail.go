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

// Snapshot usage
type CloudProjectSnapshotUsageDetail struct {

	Price OrderPrice `json:"price,omitempty"`

	// Snapshot region
	Region string `json:"region,omitempty"`

	StoredSize CloudProjectSnapshotUsageDetailStoredSize `json:"storedSize,omitempty"`
}
