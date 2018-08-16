package metrics

import (
	"time"

	"github.com/ovh/go-ovh/ovh"
)

type (
	// Client for Metrics API
	Client struct {
		ovhClient *ovh.Client
	}

	// API Request payloads

	// EditRequest edit service
	EditRequest struct {
		Description string `json:"description"`
	}

	// EditContactRequest edit contacts
	EditContactRequest struct {
		Admin   string `json:"contactAdmin,omitempty"`
		Tech    string `json:"contactTech,omitempty"`
		Billing string `json:"contactBilling,omitempty"`
	}

	// ConfirmTerminationRequest confirm termination
	ConfirmTerminationRequest struct {
		Commentary string `json:"commentary,omitempty"`
		FuturUse   string `json:"futureUse,omitempty"`
		Reason     string `json:"reason,omitempty"`
		Token      string `json:"token,omitempty"`
	}

	// ConsumptionRequest service consumption
	ConsumptionRequest struct {
		Duration int `json:"duration"`
	}

	// ConsumptionResponse service consumption
	ConsumptionResponse struct {
		MADS int `json:"mads,omitempty"`
		DDP  int `json:"ddp,omitempty"`
	}

	// LookupTokenRequest search for token
	LookupTokenRequest struct {
		Token string `json:"accessToken"`
	}

	// SetQuotaRequest set pay as you go quota
	SetQuotaRequest struct {
		MADS int `json:"quota"`
	}

	// InfosResponse billing infos
	InfosResponse struct {
		ServiceID             int64        `json:"serviceId,omitempty"`
		Domain                string       `json:"domain,omitempty"`
		Status                string       `json:"status,omitempty"`
		Renew                 ServiceRenew `json:"renew,omitempty"`
		RenewalType           string       `json:"renewalType,omitempty"`
		PossibleRenewPeriod   []int64      `json:"possibleRenewPeriod,omitempty"`
		EngagedUpTo           *time.Time   `json:"engagedUpTo,omitempty"`
		Creation              *time.Time   `json:"creation,omitempty"`
		Expiration            *time.Time   `json:"expiration,omitempty"`
		CanDeleteAtExpiration bool         `json:"canDeleteAtExpiration,omitempty"`
		ContactBilling        string       `json:"contactBilling,omitempty"`
		ContactTech           string       `json:"contactTech,omitempty"`
		ContactAdmin          string       `json:"contactAdmin,omitempty"`
	}

	// ServiceRenew billing renew config
	ServiceRenew struct {
		ManualPayment      bool  `json:"manualPayment,omitempty"`
		Period             int64 `json:"period,omitempty"`
		Forced             bool  `json:"forced,omitempty"`
		Automatic          bool  `json:"automatic,omitempty"`
		DeleteAtExpiration bool  `json:"deleteAtExpiration,omitempty"`
	}

	// EditInfosRequest set service renew options
	EditInfosRequest struct {
		Service struct {
			Renew ServiceRenew `json:"renew"`
		} `json:"service"`
	}

	// NewTokenRequest for token creation
	NewTokenRequest struct {
		Description string
		Labels      TokenLabels `json:"labels"`
		Permission  string      `json:"permission"`
	}

	// TokenLabel labelset
	TokenLabel struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	// TokenLabels list of labelset
	TokenLabels []TokenLabel

	// EditTokenRequest token options update
	EditTokenRequest struct {
		Description string `json:"description"`
	}
)
