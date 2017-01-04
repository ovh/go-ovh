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

type TelephonyEasyHuntingHuntingQueuePost struct {

	ActionOnClosure string `json:"actionOnClosure,omitempty"`

	ActionOnClosureParam string `json:"actionOnClosureParam,omitempty"`

	ActionOnOverflow string `json:"actionOnOverflow,omitempty"`

	ActionOnOverflowParam string `json:"actionOnOverflowParam,omitempty"`

	AskForRecordDisabling bool `json:"askForRecordDisabling,omitempty"`

	Description string `json:"description,omitempty"`

	MaxMember int64 `json:"maxMember,omitempty"`

	MaxWaitTime int64 `json:"maxWaitTime,omitempty"`

	Record bool `json:"record,omitempty"`

	RecordDisablingDigit int64 `json:"recordDisablingDigit,omitempty"`

	RecordDisablingLanguage string `json:"recordDisablingLanguage,omitempty"`

	SoundOnHold int64 `json:"soundOnHold,omitempty"`

	Strategy string `json:"strategy,omitempty"`
}
