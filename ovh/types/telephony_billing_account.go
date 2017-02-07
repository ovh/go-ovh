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

// TelephonyBillingAccount Billing Account
type TelephonyBillingAccount struct {
	AllowedOutplan *OrderPrice `json:"allowedOutplan,omitempty"`

	// BillingAccount Name of the billing account
	BillingAccount string `json:"billingAccount,omitempty"`

	CreditThreshold *OrderPrice `json:"creditThreshold,omitempty"`

	CurrentOutplan *OrderPrice `json:"currentOutplan,omitempty"`

	// Description Description of the billing account
	Description string `json:"description,omitempty"`

	// OverrideDisplayedNumber Override number display for calls between services of your billing account
	OverrideDisplayedNumber bool `json:"overrideDisplayedNumber,omitempty"`

	SecurityDeposit *OrderPrice `json:"securityDeposit,omitempty"`

	// Status Current status of billing account
	Status string `json:"status,omitempty"`

	// Trusted Is the billing account trusted
	Trusted bool `json:"trusted,omitempty"`
}
