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

// Detected incident
type XdslIncident struct {

	Comment string `json:"comment,omitempty"`

	// Estimated start date
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Department list
	Departments []string `json:"departments,omitempty"`

	// Estimated end date
	EndDate time.Time `json:"endDate,omitempty"`

	// ID of the incident
	Id int64 `json:"id,omitempty"`

	// NRA list
	Nra []string `json:"nra,omitempty"`

	// Operator
	Operators []string `json:"operators,omitempty"`

	// Task ID on travaux.ovh.com
	TaskId int64 `json:"taskId,omitempty"`
}
