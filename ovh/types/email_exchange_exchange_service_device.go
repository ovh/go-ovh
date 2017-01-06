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

// EmailExchangeServiceDevice Get the list of your ActiveSync devices registered on this Exchange service
type EmailExchangeServiceDevice struct {

	// IMEI International Mobile Equipment Identity
	IMEI string `json:"IMEI,omitempty"`

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// DeviceID Device Id
	DeviceID string `json:"deviceId,omitempty"`

	// DeviceModel Model device
	DeviceModel string `json:"deviceModel,omitempty"`

	// DeviceState Device State
	DeviceState string `json:"deviceState,omitempty"`

	// GUID user guid
	GUID string `json:"guid,omitempty"`

	// IDentity Exchange identity
	IDentity string `json:"identity,omitempty"`

	// LastUpdate Last update date
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// TaskPendingID Pending task id
	TaskPendingID int64 `json:"taskPendingId,omitempty"`
}