package ovh

// Domain ...
type Domain struct {
	// "Is whois obfuscation supported by this domain name's registry"
	OwoSupported bool `json:"owoSupported,omitempty"`

	// "Does the registry support ipv6 glue record"
	GlueRecordIpv6Supported bool `json:"glueRecordIpv6Supported,omitempty"`

	// "Transfer lock status"
	TransferLockStatus string `json:"transferLockStatus,omitempty"`
	//fullType: "domain.DomainLockStatusEnum"

	// "Domain's offer"
	Offer string `json:"offer,omitempty"`
	//fullType: "domain.OfferEnum"

	// "Contact Owner (you can edit it via /me/contact/<ID>)"
	WhoisOwner string `json:"whoisOwner,omitempty"`

	// "Is DNSSEC implemented for this domain name's tld"
	DnssecSupported bool `json:"dnssecSupported,omitempty"`

	// "Parent service"
	ParentService *string `json:"parentService,omitempty"`
	//fullType: "domain.ParentService"

	// "Domain name"
	Domain string `json:"domain"`

	// "Last update date"
	LastUpdate string `json:"lastUpdate,omitempty"`

	// "Does the registry support multi ip glue record"
	GlueRecordMultiIPSupported bool `json:"glueRecordMultiIpSupported,omitempty"`

	// "Name servers type"
	NameServerType string `json:"nameServerType,omitempty"`
	//fullType: "domain.DomainNsTypeEnum"
}

// DomainList list all your domain
func (c *Client) DomainList(withDetails bool) ([]Domain, error) {
	var names []string
	if err := c.Get("/domain", &names); err != nil {
		return nil, err
	}

	domains := []Domain{}
	for _, name := range names {
		domains = append(domains, Domain{Domain: name})
	}

	if !withDetails {
		return domains, nil
	}

	domainsChan, errChan := make(chan Domain), make(chan error)
	for _, domain := range domains {
		go func(domain Domain) {
			d, err := c.DomainInfo(domain.Domain)
			if err != nil {
				errChan <- err
				return
			}
			domainsChan <- *d
		}(domain)
	}

	domainsComplete := []Domain{}

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
func (c *Client) DomainInfo(domainName string) (*Domain, error) {
	domain := &Domain{}
	err := c.Get(queryEscape("/domain/%s", domainName), domain)
	return domain, err
}
