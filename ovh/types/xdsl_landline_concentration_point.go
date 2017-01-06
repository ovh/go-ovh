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

// XdslLandlineConcentrationPoint Infos about a Landline at the concentration point
type XdslLandlineConcentrationPoint struct {

	// LineHead Identifier of the head of the cable from the MDF
	LineHead string `json:"lineHead,omitempty"`

	// LineInitialSection Identifier of the section at the lineHead
	LineInitialSection int64 `json:"lineInitialSection,omitempty"`

	// LineInitialSectionPair Identifier of the pair at the lineHead's lineInitialSection
	LineInitialSectionPair int64 `json:"lineInitialSectionPair,omitempty"`
}