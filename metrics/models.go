package metrics

import "time"

type (

	// API Objects

	// Service is Metrics service
	Service struct {
		Name          string `json:"name,omitempty"`
		Description   string `json:"description,omitempty"`
		Type          string `json:"type,omitempty"`
		Region        Region `json:"region,omitempty"`
		Offer         string `json:"offer,omitempty"`
		Quota         Quota  `json:"quota,omitempty"`
		Status        string `json:"status,omitempty"`
		ShouldUpgrade bool   `json:"shouldUpgrade,omitempty"`
	}

	// Services is list of service
	Services []*Service

	// Quota of a Metrics service
	Quota struct {
		ReferenceMADS    int        `json:"ref_mads,omitempty"`
		ReferenceDDP     int        `json:"ref_ddp,omitempty"`
		Retention        int        `json:"retention,omitempty"`
		MADS             int        `json:"mads,omitempty"`
		DDP              int        `json:"ddp,omitempty"`
		Boosted          bool       `json:"boosted,omitempty"`
		LastModification *time.Time `json:"lastModification,omitempty"`
	}

	// Region of a Metrics service
	Region struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}

	// Token is a Metrics service token
	Token struct {
		ID          string      `json:"id,omitempty"`
		Access      string      `json:"access,omitempty"`
		Permission  string      `json:"permission,omitempty"`
		Description string      `json:"description,omitempty"`
		Labels      TokenLabels `json:"labels,omitempty"`
		IsRevoked   bool        `json:"isRevoked,omitempty"`
		CreatedAt   *time.Time  `json:"createdAt,omitempty"`
		ExpireAt    *time.Time  `json:"ExpireAt,omitempty"`
	}

	// Tokens is a list of token
	Tokens []*Token
)
