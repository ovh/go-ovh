package metrics

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
)

// List all Metrics services IDs
func List(c *ovh.Client) ([]string, error) {
	services := []string{}
	err := c.Get("/metrics", &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}

// Get a Metrics service with the given service name
func Get(c *ovh.Client, serviceName string) (*Service, error) {
	s := &Service{}
	err := c.Get(fmt.Sprintf("/metrics/%s", serviceName), s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Edit Metrics service options
func Edit(c *ovh.Client, serviceName string, er *EditRequest) (*Service, error) {
	s := &Service{}
	err := c.Put(fmt.Sprintf("/metrics/%s", serviceName), er, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// EditContacts update Metrics service contacts
func EditContacts(c *ovh.Client, serviceName string, ecr *EditContactRequest) error {
	return c.Post(fmt.Sprintf("/metrics/%s/changeContact", serviceName), ecr, nil)
}

// Terminate a Metrics service
func Terminate(c *ovh.Client, serviceName string) error {
	return c.Post(fmt.Sprintf("/metrics/%s/terminate", serviceName), nil, nil)
}

// ConfirmTermination must be call after service.Terminate() with the termination token (email)
func ConfirmTermination(c *ovh.Client, serviceName string, ctr *ConfirmTerminationRequest) error {
	return c.Post(fmt.Sprintf("/metrics/%s/terminate", serviceName), ctr, nil)
}

// Consumption (MADS & DDP) of the Metrics service
func Consumption(c *ovh.Client, serviceName string, cr *ConsumptionRequest) (*ConsumptionResponse, error) {
	cres := &ConsumptionResponse{}
	err := c.Post(fmt.Sprintf("/metrics/%s/consumption", serviceName), cr, cres)
	if err != nil {
		return nil, err
	}
	return cres, nil
}

// LookupToken return mathing tokens resources for a given access token
func LookupToken(c *ovh.Client, serviceName string, ltr *LookupTokenRequest) ([]string, error) {
	tokens := []string{}
	err := c.Post(fmt.Sprintf("/metrics/%s/lookup/token", serviceName), ltr, &tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

// SetQuota on a pay as you go Metrics service
func SetQuota(c *ovh.Client, serviceName string, sqr *SetQuotaRequest) error {
	return c.Put(fmt.Sprintf("/metrics/%s/quota", serviceName), sqr, nil)
}

// Infos return billing informations about Metrics service
func Infos(c *ovh.Client, serviceName string) (*InfosResponse, error) {
	ir := &InfosResponse{}
	err := c.Get(fmt.Sprintf("/metrics/%s/serviceInfos", serviceName), ir)
	if err != nil {
		return nil, err
	}
	return ir, nil
}

// EditInfos about Metrics service
func EditInfos(c *ovh.Client, serviceName string, eir *EditInfosRequest) error {
	return c.Put(fmt.Sprintf("/metrics/%s/serviceInfos", serviceName), eir, nil)
}

// ListTokens return a list of tokens ID attached to a Metrics service
func ListTokens(c *ovh.Client, serviceName string) ([]string, error) {
	tokens := []string{}
	err := c.Get(fmt.Sprintf("/metrics/%s/token", serviceName), &tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

// NewToken create a new token with the provided informations
func NewToken(c *ovh.Client, serviceName string, ntr *NewTokenRequest) (*Token, error) {
	token := &Token{}
	err := c.Post(fmt.Sprintf("/metrics/%s/token", serviceName), ntr, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// GetToken return a Metrics service token
func GetToken(c *ovh.Client, serviceName, tokenID string) (*Token, error) {
	token := &Token{}
	err := c.Get(fmt.Sprintf("/metrics/%s/token/%s", serviceName, tokenID), token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// EditToken update token options
func EditToken(c *ovh.Client, serviceName, tokenID string, etr *EditTokenRequest) (*Token, error) {
	token := &Token{}
	err := c.Put(fmt.Sprintf("/metrics/%s/token/%s", serviceName, tokenID), etr, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// RevokeToken revoke a token
func RevokeToken(c *ovh.Client, serviceName, tokenID string) error {
	return c.Delete(fmt.Sprintf("/metrics/%s/token/%s", serviceName, tokenID), nil)
}
