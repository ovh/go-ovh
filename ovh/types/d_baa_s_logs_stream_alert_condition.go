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

// Alert condition
type DBaaSLogsStreamAlertCondition struct {

	// Stream alert condition UUID
	AlertId string `json:"alertId,omitempty"`

	// Backlog size
	Backlog int64 `json:"backlog,omitempty"`

	// Alert condition type
	ConditionType string `json:"conditionType,omitempty"`

	// Constraint type
	ConstraintType string `json:"constraintType,omitempty"`

	// Field name
	Field string `json:"field,omitempty"`

	// Grace period in minutes
	Grace int64 `json:"grace,omitempty"`

	// Threshold
	Threshold int64 `json:"threshold,omitempty"`

	// Threshold condition
	ThresholdType string `json:"thresholdType,omitempty"`

	// Time lapse in minutes
	Time int64 `json:"time,omitempty"`

	// Condition label
	Title string `json:"title,omitempty"`

	// Field value
	Value string `json:"value,omitempty"`
}
