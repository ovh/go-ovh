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

// User allowed to access interfaces on your cluster
type ClusterHadoopUser struct {

	// Whether or not the User is allowed to access to the Cloudera Manager interface
	ClouderaManager bool `json:"clouderaManager,omitempty"`

	// Whether or not the User is allowed to access to the WebUI interfaces
	HttpFrontend bool `json:"httpFrontend,omitempty"`

	// Whether or not the User is allowed to access to the Hue interface
	Hue bool `json:"hue,omitempty"`

	// The username of the User
	Username string `json:"username,omitempty"`
}
