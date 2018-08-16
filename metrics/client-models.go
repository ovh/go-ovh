package metrics

import "github.com/ovh/go-ovh/ovh"

type (
	// ClientService wrap a Metrics service with client methods
	ClientService struct {
		*Service
		ovhClient *ovh.Client
	}

	// ClientServices is a list of ClientService
	ClientServices []*ClientService

	// ClientServiceToken wrap a Metrics service token with client methods
	ClientServiceToken struct {
		*Token
		Service   *Service
		ovhClient *ovh.Client
	}

	// ClientServiceTokens is a list of ClientServiceToken
	ClientServiceTokens []*ClientServiceToken
)
