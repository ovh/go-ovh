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

// Exchange mailbox information
type MsServicesExchangeInformation struct {

	// Exchange account license
	AccountLicense string `json:"accountLicense,omitempty"`

	// Indicates if the account is configured
	Configured bool `json:"configured,omitempty"`

	// Creation date
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Mailbox usage in KB
	CurrentUsage int64 `json:"currentUsage,omitempty"`

	// Delete mailbox at expiration date
	DeleteAtExpiration bool `json:"deleteAtExpiration,omitempty"`

	// Delete Outlook license at expiration date
	DeleteOutlookAtExpiration bool `json:"deleteOutlookAtExpiration,omitempty"`

	// Property needed for mailbox manual configuration (GUID)
	ExchangeGuid string `json:"exchangeGuid,omitempty"`

	// Expiration date
	ExpirationDate time.Time `json:"expirationDate,omitempty"`

	// Expiration date of Outlook license
	ExpirationOutlookDate time.Time `json:"expirationOutlookDate,omitempty"`

	// Visibility in Global Address List
	HiddenFromGAL bool `json:"hiddenFromGAL,omitempty"`

	// Account id
	Id int64 `json:"id,omitempty"`

	// Enable or disable anti-virus and anti-spam
	MailingFilter []string `json:"mailingFilter,omitempty"`

	// Outlook licence attached
	OutlookLicense bool `json:"outlookLicense,omitempty"`

	// Primary address of the mailbox
	PrimaryEmailAddress string `json:"primaryEmailAddress,omitempty"`

	// Maximum mailbox usage in GB (overall size)
	Quota int64 `json:"quota,omitempty"`

	// Frequency of Outlook license renewals
	RenewOutlookPeriod string `json:"renewOutlookPeriod,omitempty"`

	// Frequency of mailbox license renewals
	RenewPeriod string `json:"renewPeriod,omitempty"`

	SpamAndVirusConfiguration MsServicesSpamAndVirusConfiguration `json:"spamAndVirusConfiguration,omitempty"`

	// Spam activity detected on this mailbox
	SpamDetected bool `json:"spamDetected,omitempty"`

	// Ticket number of spam detection
	SpamTicketNumber int64 `json:"spamTicketNumber,omitempty"`

	// Mailbox state
	State string `json:"state,omitempty"`

	// Pending task id for this account
	TaskPendingId int64 `json:"taskPendingId,omitempty"`
}
