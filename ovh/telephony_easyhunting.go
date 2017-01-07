package ovh

import (
	"strconv"

	"github.com/runabove/go-sdk/ovh/types"
)

// TelephonyEasyHuntingList list all OVH easy calls queues associated with this billing account
// GET /telephony/{billingAccount}/easyHunting
func (c *Client) TelephonyEasyHuntingList(billingAccount string, withDetails bool) ([]types.TelephonyEasyHunting, error) {
	var names []string
	if err := c.Get(queryEscape("/telephony/%s/easyHunting", billingAccount), &names); err != nil {
		return nil, err
	}

	services := []types.TelephonyEasyHunting{}
	for _, name := range names {
		services = append(services, types.TelephonyEasyHunting{ServiceName: name})
	}

	if !withDetails {
		return services, nil
	}

	servicesChan, errChan := make(chan types.TelephonyEasyHunting), make(chan error)
	for _, telephonyEasyHunting := range services {
		go func(billingAccount, serviceName string) {
			d, err := c.TelephonyEasyHuntingInfo(billingAccount, serviceName)
			if err != nil {
				errChan <- err
				return
			}
			servicesChan <- *d
		}(billingAccount, telephonyEasyHunting.ServiceName)
	}

	servicesComplete := []types.TelephonyEasyHunting{}

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

// TelephonyEasyHuntingInfo retrieve all infos of one easy hunting service
// GET /telephony/{billingAccount}/easyHunting/{serviceName}
func (c *Client) TelephonyEasyHuntingInfo(billingAccount, serviceName string) (*types.TelephonyEasyHunting, error) {
	telephonyEasyHunting := &types.TelephonyEasyHunting{}
	err := c.Get(queryEscape("/telephony/%s/easyHunting/%s", billingAccount, serviceName), telephonyEasyHunting)
	return telephonyEasyHunting, err
}

// TelephonyOvhPabxHunting retrieves info on OVH Pabx Hunting
// GET /telephony/{billingAccount}/easyHunting/{serviceName}/hunting
func (c *Client) TelephonyOvhPabxHunting(billingAccount, serviceName string) (*types.TelephonyOvhPabxHunting, error) {
	telephonyOvhPabxHunting := &types.TelephonyOvhPabxHunting{}
	err := c.Get(queryEscape("/telephony/%s/easyHunting/%s/hunting", billingAccount, serviceName), telephonyOvhPabxHunting)
	return telephonyOvhPabxHunting, err
}

// TelephonyOvhPabxHuntingAgentList list all OVH easy calls queues associated with this billing account
// GET  /telephony/{billingAccount}/easyHunting/{serviceName}/hunting/agent
func (c *Client) TelephonyOvhPabxHuntingAgentList(billingAccount, serviceName string, withDetails bool) ([]types.TelephonyOvhPabxHuntingAgent, error) {
	var names []int64
	if err := c.Get(queryEscape("/telephony/%s/easyHunting/%s/hunting/agent", billingAccount, serviceName), &names); err != nil {
		return nil, err
	}

	agents := []types.TelephonyOvhPabxHuntingAgent{}
	for _, agentID := range names {
		agents = append(agents, types.TelephonyOvhPabxHuntingAgent{AgentID: agentID})
	}

	if !withDetails {
		return agents, nil
	}

	agentsChan, errChan := make(chan types.TelephonyOvhPabxHuntingAgent), make(chan error)
	for _, agent := range agents {
		go func(billingAccount, serviceName string, agentID int64) {
			d, err := c.TelephonyOvhPabxHuntingAgentInfo(billingAccount, serviceName, agentID)
			if err != nil {
				errChan <- err
				return
			}
			agentsChan <- *d
		}(billingAccount, serviceName, agent.AgentID)
	}

	agentsComplete := []types.TelephonyOvhPabxHuntingAgent{}

	for i := 0; i < len(agents); i++ {
		select {
		case agents := <-agentsChan:
			agentsComplete = append(agentsComplete, agents)
		case err := <-errChan:
			return nil, err
		}
	}

	return agentsComplete, nil
}

// TelephonyOvhPabxHuntingAgentInfo gets info from OVH Pabx Hunting Agent
// GET /telephony/{billingAccount}/easyHunting/{serviceName}/hunting/agent
func (c *Client) TelephonyOvhPabxHuntingAgentInfo(billingAccount, serviceName string, agentID int64) (*types.TelephonyOvhPabxHuntingAgent, error) {
	telephonyOvhPabxHuntingAgent := &types.TelephonyOvhPabxHuntingAgent{}
	err := c.Get(queryEscape("/telephony/%s/easyHunting/%s/hunting/agent/%s", billingAccount, serviceName, strconv.FormatInt(agentID, 10)), telephonyOvhPabxHuntingAgent)
	return telephonyOvhPabxHuntingAgent, err
}

// TelephonyOvhPabxHuntingAgentUpdate update OVH Pabx Hunting Agent
// PUT /telephony/{billingAccount}/easyHunting/{serviceName}/hunting/agent/{agentId}
func (c *Client) TelephonyOvhPabxHuntingAgentUpdate(billingAccount, serviceName string, agentID int64, telephonyOvhPabxHuntingAgent types.TelephonyOvhPabxHuntingAgent) (*types.TelephonyOvhPabxHuntingAgent, error) {
	r := &types.TelephonyOvhPabxHuntingAgent{}
	err := c.Put(queryEscape("/telephony/%s/easyHunting/%s/hunting/agent/%s", billingAccount, serviceName, strconv.FormatInt(agentID, 10)), telephonyOvhPabxHuntingAgent, r)
	return r, err
}
