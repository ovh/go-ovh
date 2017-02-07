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

// OrderCatalogProduct Describe a Product in the Catalog
type OrderCatalogProduct struct {

	// Configurations List of the configurations available for the product
	Configurations []*OrderCatalogConfigurationItem `json:"configurations,omitempty"`

	// Description Designation of the product
	Description string `json:"description,omitempty"`

	// Metadatas List of the metadata of the product
	Metadatas []*ComplexTypeSafeKeyValueString `json:"metadatas,omitempty"`

	// Name Plan code identifier of the product
	Name string `json:"name,omitempty"`

	// TechnicalDetails Technicals details about product
	TechnicalDetails []*ComplexTypeSafeKeyValueString `json:"technicalDetails,omitempty"`
}
