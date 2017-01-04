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

// Representation of a Association additional information
type DomainDataAssociationContact struct {

	// Contact ID related to the association contact information
	ContactId int64 `json:"contactId,omitempty"`

	// Date of the declaration of the association
	DeclarationDate time.Time `json:"declarationDate,omitempty"`

	// Association additional information ID
	Id int64 `json:"id,omitempty"`

	// Date of the publication of the declaration of the association
	PublicationDate time.Time `json:"publicationDate,omitempty"`

	// Number of the publication of the declaration of the association
	PublicationNumber string `json:"publicationNumber,omitempty"`

	// Page number of the publication of the declaration of the association
	PublicationPageNumber string `json:"publicationPageNumber,omitempty"`
}
