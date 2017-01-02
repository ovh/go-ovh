package ovh

import (
	"strconv"
)

// TelephonyEasyHunting struct
type TelephonyEasyHunting struct {
	//Max wait time when caller is in queue (in seconds)
	MaxWaitTime float64 `json:"maxWaitTime,omitempty"`

	// FeatureType
	FeatureType string `json:"featureType,omitempty"`

	// Strategy : The calls dispatching strategy
	Strategy string `json:"strategy,omitempty"`

	// QueueSize Max number of callers in queue
	QueueSize float64 `json:"queueSize,omitempty"`

	// ToneOnHold: Tone played when caller is put on hold
	ToneOnHold float64 `json:"toneOnHold,omitempty"`

	// ServiceName containers service Name
	ServiceName string `json:"serviceName,omitempty"`

	// ShowCallerNumber: The presented number when bridging calls
	ShowCallerNumber string `json:"showCallerNumber,omitempty"`

	// Description ...
	Description string `json:"description,omitempty"`

	// AnonymousRejection: Reject (hangup) anonymous calls
	AnonymousRejection bool `json:"anonymousRejection,omitempty"`

	//ToneOnOpening: Tone played when call is picked up
	ToneOnOpening float64 `json:"toneOnOpening,omitempty"`

	// serviceType
	ServiceType string `json:"serviceType,omitempty"`

	// Voicemail: The voicemail used by the EasyPABX
	Voicemail string `json:"voicemail,omitempty"`

	//ToneOnClosing: Tone played just before call is hang up
	ToneOnClosing float64 `json:"toneOnClosing,omitempty"`
}

// TelephonyOvhPabxHunting struct
type TelephonyOvhPabxHunting struct {
	// CrmUrlTemplate: The templated url of your CRM, opened by the banner application of your cloudpabx
	CRMUrlTemplate string `json:"crmUrlTemplate,omitempty"`
	// The name of your callcenter offer
	Name string `json:"name,omitempty"`
	// Enable G729 codec on your callcenter
	G729 bool `json:"g729,omitempty"`
}

// TelephonyOvhPabxHuntingAgent ...
type TelephonyOvhPabxHuntingAgent struct {
	// ID of agent
	AgentID int `json:"agentId,omitempty"`
	// The wrap up time (in seconds) after the calls
	WrapUpTime int `json:"wrapUpTime,omitempty"`
	// The number of the agent
	Number string `json:"number,omitempty"`
	// The waiting timeout (in seconds) before hangup an assigned called
	Timeout int `json:"timeout,omitempty"`
	// The current status of the agent
	Status string `json:"status,omitempty"`
	// The maximum of simultaneous calls that the agent will receive from the hunting
	SimultaneousLines int `json:"simultaneousLines,omitempty"`
	// The id of the current break status of the agent
	BreakStatus int `json:"breakStatus,omitempty"`
}

// TelephonyEasyHuntingList list all OVH easy calls queues associated with this billing account
// GET /telephony/{billingAccount}/easyHunting
func (c *Client) TelephonyEasyHuntingList(billingAccount string, withDetails bool) ([]TelephonyEasyHunting, error) {
	var names []string
	if err := c.Get(queryEscape("/telephony/%s/easyHunting", billingAccount), &names); err != nil {
		return nil, err
	}

	services := []TelephonyEasyHunting{}
	for _, name := range names {
		services = append(services, TelephonyEasyHunting{ServiceName: name})
	}

	if !withDetails {
		return services, nil
	}

	servicesChan, errChan := make(chan TelephonyEasyHunting), make(chan error)
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

	servicesComplete := []TelephonyEasyHunting{}

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
func (c *Client) TelephonyEasyHuntingInfo(billingAccount, serviceName string) (*TelephonyEasyHunting, error) {
	telephonyEasyHunting := &TelephonyEasyHunting{}
	err := c.Get(queryEscape("/telephony/%s/easyHunting/%s", billingAccount, serviceName), telephonyEasyHunting)
	return telephonyEasyHunting, err
}

// TelephonyOvhPabxHunting retrieves info on OVH Pabx Hunting
// GET /telephony/{billingAccount}/easyHunting/{serviceName}/hunting
func (c *Client) TelephonyOvhPabxHunting(billingAccount, serviceName string) (*TelephonyOvhPabxHunting, error) {
	telephonyOvhPabxHunting := &TelephonyOvhPabxHunting{}
	err := c.Get(queryEscape("/telephony/%s/easyHunting/%s/hunting", billingAccount, serviceName), telephonyOvhPabxHunting)
	return telephonyOvhPabxHunting, err
}

// TelephonyOvhPabxHuntingAgentList list all OVH easy calls queues associated with this billing account
// GET  /telephony/{billingAccount}/easyHunting/{serviceName}/hunting/agent
func (c *Client) TelephonyOvhPabxHuntingAgentList(billingAccount, serviceName string, withDetails bool) ([]TelephonyOvhPabxHuntingAgent, error) {
	var names []int
	if err := c.Get(queryEscape("/telephony/%s/easyHunting/%s/hunting/agent", billingAccount, serviceName), &names); err != nil {
		return nil, err
	}

	agents := []TelephonyOvhPabxHuntingAgent{}
	for _, agentID := range names {
		agents = append(agents, TelephonyOvhPabxHuntingAgent{AgentID: agentID})
	}

	if !withDetails {
		return agents, nil
	}

	agentsChan, errChan := make(chan TelephonyOvhPabxHuntingAgent), make(chan error)
	for _, agent := range agents {
		go func(billingAccount, serviceName string, agentID int) {
			d, err := c.TelephonyOvhPabxHuntingAgentInfo(billingAccount, serviceName, agentID)
			if err != nil {
				errChan <- err
				return
			}
			agentsChan <- *d
		}(billingAccount, serviceName, agent.AgentID)
	}

	agentsComplete := []TelephonyOvhPabxHuntingAgent{}

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
func (c *Client) TelephonyOvhPabxHuntingAgentInfo(billingAccount, serviceName string, agentID int) (*TelephonyOvhPabxHuntingAgent, error) {
	telephonyOvhPabxHuntingAgent := &TelephonyOvhPabxHuntingAgent{}
	err := c.Get(queryEscape("/telephony/%s/easyHunting/%s/hunting/agent/%s", billingAccount, serviceName, strconv.Itoa(agentID)), telephonyOvhPabxHuntingAgent)
	return telephonyOvhPabxHuntingAgent, err
}

// TelephonyOvhPabxHuntingAgentUpdate update OVH Pabx Hunting Agent
// PUT /telephony/{billingAccount}/easyHunting/{serviceName}/hunting/agent/{agentId}
func (c *Client) TelephonyOvhPabxHuntingAgentUpdate(billingAccount, serviceName string, agentID int, telephonyOvhPabxHuntingAgent TelephonyOvhPabxHuntingAgent) (*TelephonyOvhPabxHuntingAgent, error) {
	r := &TelephonyOvhPabxHuntingAgent{}
	err := c.Put(queryEscape("/telephony/%s/easyHunting/%s/hunting/agent/%s", billingAccount, serviceName, strconv.Itoa(agentID)), telephonyOvhPabxHuntingAgent, r)
	return r, err
}
