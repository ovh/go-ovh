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

// PaasDatabaseInstanceParameters Parameters used for the instance creation
type PaasDatabaseInstanceParameters struct {

	// ImageName Image used in the new instance
	ImageName string `json:"imageName,omitempty"`

	// Name Custom name
	Name string `json:"name,omitempty"`

	// OfferName Offer associated with this new instance
	OfferName string `json:"offerName,omitempty"`

	// RegionName Region where you want to run the new instance
	RegionName string `json:"regionName,omitempty"`
}
