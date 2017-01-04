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

// A structure describing the current password policy for your Dedicated Cloud
type DedicatedCloudPasswordPolicy struct {

	// List of denied characters in the password
	DeniedChars []string `json:"deniedChars,omitempty"`

	// Whether or not a digit (0-9) is mandatory in the password
	DigitMandatory bool `json:"digitMandatory,omitempty"`

	// Whether or not a letter (a-z or A-Z) is mandatory in the password
	LetterMandatory bool `json:"letterMandatory,omitempty"`

	// Whether or not a lowercase letter (a-z) is mandatory in the password
	LowercaseLetterMandatory bool `json:"lowercaseLetterMandatory,omitempty"`

	// Maximum lenght of the password
	MaxLength int64 `json:"maxLength,omitempty"`

	// Minimum lenght of the password
	MinLength int64 `json:"minLength,omitempty"`

	// Whether or not a special character (\\W or _) is mandatory in the password
	SpecialMandatory bool `json:"specialMandatory,omitempty"`

	// Whether or not an uppercase letter (A-Z) is mandatory in the password
	UppercaseLetterMandatory bool `json:"uppercaseLetterMandatory,omitempty"`
}
