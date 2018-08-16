package metrics

import (
	"github.com/ovh/go-ovh/ovh"
)

// NewClient for Metrics API based on OVH API one
func NewClient(ovhClient *ovh.Client) *Client {
	return &Client{
		ovhClient: ovhClient,
	}
}

// Service related calls

// Services return a service list
func (c *Client) Services() (ClientServices, error) {
	serviceNames, err := List(c.ovhClient)
	if err != nil {
		return nil, err
	}

	services := make(ClientServices, len(serviceNames))
	for i, serviceName := range serviceNames {
		service, err := Get(c.ovhClient, serviceName)
		if err != nil {
			return nil, err
		}
		if service == nil {
			continue
		}

		services[i] = &ClientService{
			Service:   service,
			ovhClient: c.ovhClient,
		}
	}

	return services, nil
}

// Service return a Metrics service
func (c *Client) Service(serviceName string) (*ClientService, error) {
	service, err := Get(c.ovhClient, serviceName)
	if err != nil {
		return nil, err
	}

	return &ClientService{
		Service:   service,
		ovhClient: c.ovhClient,
	}, nil
}

// Edit Metrics service
func (cs *ClientService) Edit(er *EditRequest) error {
	service, err := Edit(cs.ovhClient, cs.Name, er)
	if err != nil {
		return err
	}
	cs = &ClientService{
		Service: service,
	}
	return nil
}

// EditContacts Metrics service
func (cs *ClientService) EditContacts(ec *EditContactRequest) error {
	return EditContacts(cs.ovhClient, cs.Name, ec)
}

// Terminate Metrics service
func (cs *ClientService) Terminate() error {
	return Terminate(cs.ovhClient, cs.Name)
}

// ConfirmTermination confirm a Metrics service termination
func (cs *ClientService) ConfirmTermination(ct *ConfirmTerminationRequest) error {
	return ConfirmTermination(cs.ovhClient, cs.Name, ct)
}

// Tokens return a list of tokens for a service
func (cs *ClientService) Tokens() (ClientServiceTokens, error) {
	tokenIDs, err := ListTokens(cs.ovhClient, cs.Name)
	if err != nil {
		return nil, err
	}

	tokens := make(ClientServiceTokens, len(tokenIDs))
	for i, tokenID := range tokenIDs {
		token, err := GetToken(cs.ovhClient, cs.Name, tokenID)
		if err != nil {
			return nil, err
		}
		if token == nil {
			continue
		}

		tokens[i] = &ClientServiceToken{
			Token:     token,
			Service:   cs.Service,
			ovhClient: cs.ovhClient,
		}
	}

	return tokens, nil
}

// Consumption return the current service consumption
func (cs *ClientService) Consumption(cr *ConsumptionRequest) (*ConsumptionResponse, error) {
	return Consumption(cs.ovhClient, cs.Name, cr)
}

// LookupToken return the mathing tokens
func (cs *ClientService) LookupToken(ltr *LookupTokenRequest) (ClientServiceTokens, error) {
	tokenIDs, err := LookupToken(cs.ovhClient, cs.Name, ltr)
	if err != nil {
		return nil, err
	}

	tokens := make(ClientServiceTokens, len(tokenIDs))
	for i, tokenID := range tokenIDs {
		token, err := GetToken(cs.ovhClient, cs.Name, tokenID)
		if err != nil {
			return nil, err
		}
		if token == nil {
			continue
		}

		tokens[i] = &ClientServiceToken{
			Token:     token,
			Service:   cs.Service,
			ovhClient: cs.ovhClient,
		}
	}

	return tokens, nil
}

// SetQuota Set overquota limit on Metrics service
func (cs *ClientService) SetQuota(sqr *SetQuotaRequest) error {
	if err := SetQuota(cs.ovhClient, cs.Name, sqr); err != nil {
		return err
	}
	cs.Quota.MADS = sqr.MADS
	return nil
}

// Infos return the current service billing infos
func (cs *ClientService) Infos() (*InfosResponse, error) {
	return Infos(cs.ovhClient, cs.Name)
}

// EditInfos update Metrics service billing infos
func (cs *ClientService) EditInfos(eir *EditInfosRequest) error {
	return EditInfos(cs.ovhClient, cs.Name, eir)
}

// NewToken update Metrics service billing infos
func (cs *ClientService) NewToken(ntr *NewTokenRequest) (*ClientServiceToken, error) {
	token, err := NewToken(cs.ovhClient, cs.Name, ntr)
	if err != nil {
		return nil, err
	}

	return &ClientServiceToken{
		Token:     token,
		Service:   cs.Service,
		ovhClient: cs.ovhClient,
	}, nil
}

// Token return the token on Metrics service with the provided token ID
func (cs *ClientService) Token(tokenID string) (*ClientServiceToken, error) {
	token, err := GetToken(cs.ovhClient, cs.Name, tokenID)
	if err != nil {
		return nil, err
	}

	return &ClientServiceToken{
		Token:     token,
		ovhClient: cs.ovhClient,
	}, nil
}

// Tokens related calls

// Edit upddate token options
func (ct *ClientServiceToken) Edit(etr *EditTokenRequest) error {
	token, err := EditToken(ct.ovhClient, ct.Service.Name, ct.ID, etr)
	if err != nil {
		return err
	}

	ct.Token = token
	return nil
}

// Revoke a token
func (ct *ClientServiceToken) Revoke() error {
	return RevokeToken(ct.ovhClient, ct.Service.Name, ct.ID)
}
