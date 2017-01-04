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

type MeContactPost struct {

	Address ContactAddress `json:"address,omitempty"`

	BirthCity string `json:"birthCity,omitempty"`

	BirthCountry string `json:"birthCountry,omitempty"`

	BirthDay time.Time `json:"birthDay,omitempty"`

	BirthZip string `json:"birthZip,omitempty"`

	CellPhone string `json:"cellPhone,omitempty"`

	CompanyNationalIdentificationNumber string `json:"companyNationalIdentificationNumber,omitempty"`

	Email string `json:"email,omitempty"`

	Fax string `json:"fax,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Gender string `json:"gender,omitempty"`

	Language string `json:"language,omitempty"`

	LastName string `json:"lastName,omitempty"`

	LegalForm string `json:"legalForm,omitempty"`

	NationalIdentificationNumber string `json:"nationalIdentificationNumber,omitempty"`

	Nationality string `json:"nationality,omitempty"`

	OrganisationName string `json:"organisationName,omitempty"`

	OrganisationType string `json:"organisationType,omitempty"`

	Phone string `json:"phone,omitempty"`

	Vat string `json:"vat,omitempty"`
}
