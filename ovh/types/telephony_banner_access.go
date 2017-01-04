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

// The web access for your cloudpabx
type TelephonyBannerAccess struct {

	// The creation date of this access
	CreationDate time.Time `json:"creationDate,omitempty"`

	Id int64 `json:"id,omitempty"`

	// The url of the web access
	Url string `json:"url,omitempty"`
}
