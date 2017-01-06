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

// TelephonyConferenceParticipants Conference service
type TelephonyConferenceParticipants struct {

	// CallerName The caller name of the participant
	CallerName string `json:"callerName,omitempty"`

	// CallerNumber The caller number of the participant
	CallerNumber string `json:"callerNumber,omitempty"`

	// Energy The current level of the participant audio transmission
	Energy int64 `json:"energy,omitempty"`

	// Floor Whether or not the participant is active in the room
	Floor bool `json:"floor,omitempty"`

	// Hear Whether or not the participant can hear the conference
	Hear bool `json:"hear,omitempty"`

	// ID The id of the participant
	ID int64 `json:"id,omitempty"`

	// Speak Whether or not the participant can talk in the conference
	Speak bool `json:"speak,omitempty"`

	// Talking Whether or not the participant is talking
	Talking bool `json:"talking,omitempty"`
}