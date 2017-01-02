package ovh

// Telephony struct
type Telephony struct {
	// SecurityDeposit contains Security deposit amount
	SecurityDeposit string `json:"securityDeposit,omitempty"`
	// Status contains Current status of billing account
	Status string `json:"status,omitempty"`
	//  OverrideDisplayedNumber contains Override number display for calls between services of your billing account
	OverrideDisplayedNumber bool `json:"overrideDisplayedNumber,omitempty"`
	// CurrentOutplan contains Price with it's currency and textual representation
	CurrentOutplan string `json:"currentOutplan,omitempty"`
	// Trusted : Is the billing account trusted
	Trusted bool `json:"trusted,omitempty"`
	// Description of the billing account
	Description string `json:"description,omitempty"`
	// AllowedOutplan Allowed outplan
	AllowedOutplan string `json:"allowedOutplan,omitempty"`
	// BillingAccount : Name of the billing account
	BillingAccount string `json:"billingAccount,omitempty"`
	// CreditThreshold : Allowed threshold credit
	CreditThreshold string `json:"creditThreshold,omitempty"`
}

// TelephonyListBillingAccount list all your telephony services
func (c *Client) TelephonyListBillingAccount(withDetails bool) ([]Telephony, error) {
	var names []string
	if err := c.Get("/telephony", &names); err != nil {
		return nil, err
	}

	services := []Telephony{}
	for _, name := range names {
		services = append(services, Telephony{BillingAccount: name})
	}

	if !withDetails {
		return services, nil
	}

	servicesChan, errChan := make(chan Telephony), make(chan error)
	for _, telephony := range services {
		go func(telephony Telephony) {
			d, err := c.TelephonyBillingAccountInfo(telephony.BillingAccount)
			if err != nil {
				errChan <- err
				return
			}
			servicesChan <- *d
		}(telephony)
	}

	servicesComplete := []Telephony{}

	for i := 0; i < len(services); i++ {
		select {
		case services := <-servicesChan:
			servicesComplete = append(servicesComplete, services)
		case err := <-errChan:
			return nil, err
		}
	}

	return servicesComplete, nil
}

// TelephonyBillingAccountInfo retrieve all infos of one of your services
func (c *Client) TelephonyBillingAccountInfo(billingAccount string) (*Telephony, error) {
	telephony := &Telephony{}
	err := c.Get(queryEscape("/telephony/%s", billingAccount), telephony)
	return telephony, err
}
