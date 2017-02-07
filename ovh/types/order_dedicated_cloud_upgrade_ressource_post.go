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

// OrderDedicatedCloudUpgradeRessourcePost ...
type OrderDedicatedCloudUpgradeRessourcePost struct {
	UpgradeType string `json:"upgradeType,omitempty"`

	UpgradedRessourceID int64 `json:"upgradedRessourceId,omitempty"`

	UpgradedRessourceType string `json:"upgradedRessourceType,omitempty"`
}
