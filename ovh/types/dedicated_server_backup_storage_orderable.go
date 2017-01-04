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

// A structure describing informations storage backup orderable for this dedicated server
type DedicatedServerBackupStorageOrderable struct {

	// Backup storage orderable capacities in gigabytes
	Capacities []int64 `json:"capacities,omitempty"`

	// Is a backup storage is orderable for this server
	Orderable bool `json:"orderable,omitempty"`
}
