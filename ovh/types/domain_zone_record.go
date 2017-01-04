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

// Zone resource records
type DomainZoneRecord struct {

	// Resource record Name
	FieldType string `json:"fieldType,omitempty"`

	// Id of the zone resource record
	Id int64 `json:"id,omitempty"`

	// Resource record subdomain
	SubDomain string `json:"subDomain,omitempty"`

	// Resource record target
	Target string `json:"target,omitempty"`

	// Resource record ttl
	Ttl int64 `json:"ttl,omitempty"`

	// Resource record zone
	Zone string `json:"zone,omitempty"`
}
