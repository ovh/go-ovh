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

type DedicatedServerServiceMonitoringPost struct {

	ChallengeText string `json:"challengeText,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Interval string `json:"interval,omitempty"`

	Ip string `json:"ip,omitempty"`

	Port int64 `json:"port,omitempty"`

	Protocol string `json:"protocol,omitempty"`

	Url string `json:"url,omitempty"`
}
