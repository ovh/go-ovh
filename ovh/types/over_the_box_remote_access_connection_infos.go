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

// OverTheBoxRemoteAccessConnectionInfos All the infos needed to connect yourself to your OTB
type OverTheBoxRemoteAccessConnectionInfos struct {

	// IP IP to connect to when accessing the device remotely
	IP string `json:"ip,omitempty"`

	// Port Port to connect to when accessing the device remotely
	Port int64 `json:"port,omitempty"`
}
