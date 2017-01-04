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

// Home Location Register informations. Give informations about a given cellular phone.
type SmsHlrLookupNumber struct {

	// HLR creation datetime
	Datetime time.Time `json:"datetime,omitempty"`

	// HLR id
	Id int64 `json:"id,omitempty"`

	// MSISDN
	Msisdn string `json:"msisdn,omitempty"`

	// The {Mobile Country Code, Mobile Network Code} unique identifier
	OperatorCode string `json:"operatorCode,omitempty"`

	// Has the MSISDN been ported from its original network
	Ported bool `json:"ported,omitempty"`

	// Is the MSISDN currently reachable
	Reachable bool `json:"reachable,omitempty"`

	// Is the MSISDN currently roaming outside its natinal network
	Roaming bool `json:"roaming,omitempty"`

	// Status of the HLR request
	Status string `json:"status,omitempty"`

	// Is the MSISDN valid
	Valid bool `json:"valid,omitempty"`
}
