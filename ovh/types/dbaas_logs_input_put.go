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

type DbaasLogsInputPut struct {

	Description string `json:"description,omitempty"`

	EngineId string `json:"engineId,omitempty"`

	ExposedPort string `json:"exposedPort,omitempty"`

	OptionId string `json:"optionId,omitempty"`

	SingleInstance bool `json:"singleInstance,omitempty"`

	StreamId string `json:"streamId,omitempty"`

	Title string `json:"title,omitempty"`
}
