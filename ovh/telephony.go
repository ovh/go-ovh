package ovh

import (
	"github.com/runabove/go-sdk/ovh/types"
)

// TelephonyListBillingAccount list all your telephony services
func (c *Client) TelephonyListBillingAccount(withDetails bool) ([]types.TelephonyBillingAccount, error) {
	var names []string
	if err := c.Get("/telephony", &names); err != nil {
		return nil, err
	}

	services := []types.TelephonyBillingAccount{}
	for _, name := range names {
		services = append(services, types.TelephonyBillingAccount{BillingAccount: name})
	}

	if !withDetails {
		return services, nil
	}

	servicesChan, errChan := make(chan types.TelephonyBillingAccount), make(chan error)
	for _, telephony := range services {
		go func(telephony types.TelephonyBillingAccount) {
			d, err := c.TelephonyBillingAccountInfo(telephony.BillingAccount)
			if err != nil {
				errChan <- err
				return
			}
			servicesChan <- *d
		}(telephony)
	}

	servicesComplete := []types.TelephonyBillingAccount{}

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
func (c *Client) TelephonyBillingAccountInfo(billingAccount string) (*types.TelephonyBillingAccount, error) {
	telephony := &types.TelephonyBillingAccount{}
	err := c.Get(queryEscape("/telephony/%s", billingAccount), telephony)
	return telephony, err
}
