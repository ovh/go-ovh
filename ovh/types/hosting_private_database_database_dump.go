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

// Dump
type HostingPrivateDatabaseDatabaseDump struct {

	// Creation date of the dump
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Dump of this database name
	DatabaseName string `json:"databaseName,omitempty"`

	// Automatic deletion date of the dump
	DeletionDate time.Time `json:"deletionDate,omitempty"`

	// Dump id
	Id int64 `json:"id,omitempty"`

	// Dump url access
	Url string `json:"url,omitempty"`
}
