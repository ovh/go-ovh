package ovh

import (
	"github.com/runabove/go-sdk/ovh/types"
)

// DomainList list all your domain
func (c *Client) DomainList(withDetails bool) ([]types.DomainDomain, error) {
	var names []string
	if err := c.Get("/domain", &names); err != nil {
		return nil, err
	}

	domains := []types.DomainDomain{}
	for _, name := range names {
		domains = append(domains, types.DomainDomain{Domain: name})
	}

	if !withDetails {
		return domains, nil
	}

	domainsChan, errChan := make(chan types.DomainDomain), make(chan error)
	for _, domain := range domains {
		go func(domain types.DomainDomain) {
			d, err := c.DomainInfo(domain.Domain)
			if err != nil {
				errChan <- err
				return
			}
			domainsChan <- *d
		}(domain)
	}

	domainsComplete := []types.DomainDomain{}

	for i := 0; i < len(domains); i++ {
		select {
		case domains := <-domainsChan:
			domainsComplete = append(domainsComplete, domains)
		case err := <-errChan:
			return nil, err
		}
	}

	return domainsComplete, nil
}

// DomainInfo retrieve all infos of one of your domains
func (c *Client) DomainInfo(domainName string) (*types.DomainDomain, error) {
	domain := &types.DomainDomain{}
	err := c.Get(queryEscape("/domain/%s", domainName), domain)
	return domain, err
}
