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

type MePaymentMeanPaypalPost struct {

	Description string `json:"description,omitempty"`

	ReturnUrl string `json:"returnUrl,omitempty"`

	SetDefault bool `json:"setDefault,omitempty"`
}
