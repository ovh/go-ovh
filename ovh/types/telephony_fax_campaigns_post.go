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

type TelephonyFaxCampaignsPost struct {

	DocumentId string `json:"documentId,omitempty"`

	Name string `json:"name,omitempty"`

	RecipientsDocId string `json:"recipientsDocId,omitempty"`

	RecipientsList []string `json:"recipientsList,omitempty"`

	RecipientsType string `json:"recipientsType,omitempty"`

	SendDate time.Time `json:"sendDate,omitempty"`

	SendType string `json:"sendType,omitempty"`
}
