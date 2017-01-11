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

// Domain Domain name administration
type Domain struct {

	// DNSsecSupported Is DNSSEC implemented for this domain name's tld
	DNSsecSupported bool `json:"dnssecSupported,omitempty"`

	// Domain Domain name
	Domain string `json:"domain,omitempty"`

	// GlueRecordIPv6Supported Does the registry support ipv6 glue record
	GlueRecordIPv6Supported bool `json:"glueRecordIpv6Supported,omitempty"`

	// GlueRecordMultiIPSupported Does the registry support multi ip glue record
	GlueRecordMultiIPSupported bool `json:"glueRecordMultiIpSupported,omitempty"`

	// LastUpdate Last update date
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// NameServerType Name servers type
	NameServerType string `json:"nameServerType,omitempty"`

	// Offer Domain's offer
	Offer string `json:"offer,omitempty"`

	// OwoSupported Is whois obfuscation supported by this domain name's registry
	OwoSupported bool `json:"owoSupported,omitempty"`

	ParentService *DomainParentService `json:"parentService,omitempty"`

	// TransferLockStatus Transfer lock status
	TransferLockStatus string `json:"transferLockStatus,omitempty"`

	// WhoisOwner Contact Owner (you can edit it via /me/contact/<ID>)
	WhoisOwner string `json:"whoisOwner,omitempty"`
}
