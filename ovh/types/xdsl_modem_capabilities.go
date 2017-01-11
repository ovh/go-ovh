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

// XdslModemCapabilities Describe the capabilities of the Modem
type XdslModemCapabilities struct {
	CanBeManagedByOvh bool `json:"canBeManagedByOvh,omitempty"`

	CanChangeBridgeMode bool `json:"canChangeBridgeMode,omitempty"`

	CanChangeDHCP bool `json:"canChangeDHCP,omitempty"`

	CanChangeDMZ bool `json:"canChangeDMZ,omitempty"`

	CanChangeEasyFirewallLevel bool `json:"canChangeEasyFirewallLevel,omitempty"`

	CanChangeLAN bool `json:"canChangeLAN,omitempty"`

	CanChangeManagement bool `json:"canChangeManagement,omitempty"`

	CanChangeMtu bool `json:"canChangeMtu,omitempty"`

	CanChangePortMapping bool `json:"canChangePortMapping,omitempty"`

	CanChangeWLAN bool `json:"canChangeWLAN,omitempty"`

	CanReboot bool `json:"canReboot,omitempty"`

	CanRefreshConnectedDevices bool `json:"canRefreshConnectedDevices,omitempty"`

	CanReset bool `json:"canReset,omitempty"`
}
