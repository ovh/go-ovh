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

// Time condition
type TelephonyOvhPabxDialplanExtensionConditionTime struct {

	ConditionId int64 `json:"conditionId,omitempty"`

	// The time of the day when the extension will start to be executed
	TimeFrom string `json:"timeFrom,omitempty"`

	// The time of the day when the extension will stop to be executed
	TimeTo string `json:"timeTo,omitempty"`

	// The day of the week when the extension will be executed
	WeekDay string `json:"weekDay,omitempty"`
}
