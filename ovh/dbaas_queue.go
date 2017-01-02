package ovh

import (
	"fmt"
)

// DBaasQueueApp represents a Qaas application
type DBaasQueueApp struct {
	ID        string `json:"id,omitempty"  description:"Application ID"`
	HumanID   string `json:"humanId,omitempty"  description:"Human ID of the application"`
	Name      string `json:"name,omitempty"  description:"Application name"`
	RegionID  string `json:"regionId,omitempty"  description:"Region ID of the application"`
	AppStatus string `json:"status,omitempty"  description:"Application status" enum:"not_configured,active,deleted"`
}

// DBaasQueueKey represents a Qaas key used to authenticate users
type DBaasQueueKey struct {
	ID         string `json:"id" description:"Key ID"`
	Name       string `json:"name" description:"Key name"`
	HumanAppID string `json:"humanAppId" description:"Human ID of the key's application"`
}

// DBaasQueueRole represents a Qaas Role
type DBaasQueueRole struct {
	Name          string   `json:"name" description:"Role name"`
	ReadACL       []string `json:"readAcl" description:"List of topics with read access"`
	WriteACL      []string `json:"writeAcl" description:"List of topics with write access"`
	AutoCreateACL bool     `json:"autoCreateACL" description:"Automatically create non-existing topics on read & write operations"`
}

// DBaasQueueRegion represents a Qaas region
type DBaasQueueRegion struct {
	ID   string `json:"id" description:"Region ID"`
	Name string `json:"name" description:"Region name"`
	URL  string `json:"url" description:"Region URL"`
}

// DBaasQueueTopic represents a Kafka topic
type DBaasQueueTopic struct {
	ID                string `json:"id" description:"Topic ID" schema:"string"`
	Partitions        int    `json:"partitions" description:"Number of partitions"`
	ReplicationFactor int    `json:"replicationFactor" description:"Replication factor"`
}

// DBaasQueueUser represents a Qaas user
type DBaasQueueUser struct {
	ID       string   `json:"id" description:"User ID" schema:"string"`
	Name     string   `json:"name" description:"User name"`
	Roles    []string `json:"roles" description:"List of roles this user belongs to"` // not stored like that in DB
	Password string   `json:"password,omitempty" description:"User Password" schema:"string"`
}

// DBaasQueueMetricsAccount represents metrics account
type DBaasQueueMetricsAccount struct {
	Host  string `json:"host" description:"OpenTSDB host url"`
	Token string `json:"token" description:"Token for OpenTSDB metrics access"`
}

// DBaasQueueServiceInfo contains info about a service
type DBaasQueueServiceInfo struct {
	CanDeleteAtExpiration bool        `json:"canDeleteAtExpiration"`
	ContactAdmin          string      `json:"contactAdmin"`
	ContactBilling        string      `json:"contactBilling"`
	ContactTech           string      `json:"contactTech"`
	Creation              string      `json:"creation"`
	Domain                string      `json:"domain"`
	EngagedUpTo           interface{} `json:"engagedUpTo"`
	Expiration            string      `json:"expiration"`
	PossibleRenewPeriod   []int       `json:"possibleRenewPeriod"`
	Renew                 struct {
		Automatic          bool `json:"automatic"`
		DeleteAtExpiration bool `json:"deleteAtExpiration"`
		Forced             bool `json:"forced"`
		Period             int  `json:"period"`
	} `json:"renew"`
	RenewalType string `json:"renewalType"`
	Status      string `json:"status"`
}

// DBaasQueueAppList list all your app
func (c *Client) DBaasQueueAppList(withDetails bool) ([]DBaasQueueApp, error) {
	var ids []string
	if err := c.Get("/dbaas/queue", &ids); err != nil {
		return nil, err
	}

	apps := []DBaasQueueApp{}
	for _, name := range ids {
		apps = append(apps, DBaasQueueApp{ID: name})
	}

	if !withDetails {
		return apps, nil
	}

	appsChan, errChan := make(chan DBaasQueueApp), make(chan error)
	for _, app := range apps {
		go func(app DBaasQueueApp) {
			d, err := c.DBaasQueueAppInfo(app.ID)
			if err != nil {
				errChan <- err
				return
			}
			appsChan <- *d
		}(app)
	}

	appsComplete := []DBaasQueueApp{}
	for i := 0; i < len(apps); i++ {
		select {
		case apps := <-appsChan:
			appsComplete = append(appsComplete, apps)
		case err := <-errChan:
			return nil, err
		}
	}

	return appsComplete, nil
}

// DBaasQueueAppInfo retrieve all infos of one of your apps
func (c *Client) DBaasQueueAppInfo(serviceName string) (*DBaasQueueApp, error) {
	app := &DBaasQueueApp{}
	err := c.Get(queryEscape("/dbaas/queue/%s", serviceName), app)
	return app, err
}

// DBaasQueueAppServiceInfo retrieve all infos of one of your apps
func (c *Client) DBaasQueueAppServiceInfo(serviceName string) (*DBaasQueueServiceInfo, error) {
	app := &DBaasQueueServiceInfo{}
	err := c.Get(queryEscape("/dbaas/queue/%s/serviceInfos", serviceName), app)
	return app, err
}

// DBaasQueueAppInfoByName retrieve all infos of one of your apps
func (c *Client) DBaasQueueAppInfoByName(name string) (*DBaasQueueApp, error) {
	apps, err := c.DBaasQueueAppList(true)
	if err != nil {
		return nil, err
	}
	for _, app := range apps {
		if app.Name == name {
			return &app, nil
		}
	}

	return nil, fmt.Errorf("No App found with name:%s", name)
}

// DBaasQueueKeyList list all key on a service
func (c *Client) DBaasQueueKeyList(serviceName string, withDetails bool) ([]DBaasQueueKey, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/key", serviceName), &ids); err != nil {
		return nil, err
	}

	keys := []DBaasQueueKey{}
	for _, name := range ids {
		keys = append(keys, DBaasQueueKey{ID: name})
	}

	if !withDetails {
		return keys, nil
	}

	keysChan, errChan := make(chan DBaasQueueKey), make(chan error)
	for _, key := range keys {
		go func(key DBaasQueueKey) {
			d, err := c.DBaasQueueKeyInfo(serviceName, key.ID)
			if err != nil {
				errChan <- err
				return
			}
			keysChan <- *d
		}(key)
	}

	keysComplete := []DBaasQueueKey{}
	for i := 0; i < len(keys); i++ {
		select {
		case keys := <-keysChan:
			keysComplete = append(keysComplete, keys)
		case err := <-errChan:
			return nil, err
		}
	}

	return keysComplete, nil
}

// DBaasQueueKeyInfo retrieves all infos of one of your apps
func (c *Client) DBaasQueueKeyInfo(serviceName, keyID string) (*DBaasQueueKey, error) {
	key := &DBaasQueueKey{}
	err := c.Get(queryEscape("/dbaas/queue/%s/key/%s", serviceName, keyID), key)
	return key, err
}

// DBaasQueueRoleList list all roles on a service
func (c *Client) DBaasQueueRoleList(serviceName string, withDetails bool) ([]DBaasQueueRole, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/role", serviceName), &ids); err != nil {
		return nil, err
	}

	roles := []DBaasQueueRole{}
	for _, name := range ids {
		roles = append(roles, DBaasQueueRole{Name: name})
	}

	if !withDetails {
		return roles, nil
	}

	rolesChan, errChan := make(chan DBaasQueueRole), make(chan error)
	for _, role := range roles {
		go func(role DBaasQueueRole) {
			d, err := c.DBaasQueueRoleInfo(serviceName, role.Name)
			if err != nil {
				errChan <- err
				return
			}
			rolesChan <- *d
		}(role)
	}

	rolesComplete := []DBaasQueueRole{}
	for i := 0; i < len(roles); i++ {
		select {
		case roles := <-rolesChan:
			rolesComplete = append(rolesComplete, roles)
		case err := <-errChan:
			return nil, err
		}
	}

	return rolesComplete, nil
}

// DBaasQueueRoleInfo  retrieves all infos of one role on a service
func (c *Client) DBaasQueueRoleInfo(serviceName, roleID string) (*DBaasQueueRole, error) {
	role := &DBaasQueueRole{}
	err := c.Get(queryEscape("/dbaas/queue/%s/role/%s", serviceName, roleID), role)
	return role, err
}

// DBaasQueueRegionList list all region on a service
func (c *Client) DBaasQueueRegionList(serviceName string, withDetails bool) ([]DBaasQueueRegion, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/region", serviceName), &ids); err != nil {
		return nil, err
	}

	regions := []DBaasQueueRegion{}
	for _, name := range ids {
		regions = append(regions, DBaasQueueRegion{Name: name})
	}

	if !withDetails {
		return regions, nil
	}

	regionsChan, errChan := make(chan DBaasQueueRegion), make(chan error)
	for _, region := range regions {
		go func(region DBaasQueueRegion) {
			d, err := c.DBaasQueueRegionInfo(serviceName, region.Name)
			if err != nil {
				errChan <- err
				return
			}
			regionsChan <- *d
		}(region)
	}

	regionsComplete := []DBaasQueueRegion{}
	for i := 0; i < len(regions); i++ {
		select {
		case regions := <-regionsChan:
			regionsComplete = append(regionsComplete, regions)
		case err := <-errChan:
			return nil, err
		}
	}

	return regionsComplete, nil
}

// DBaasQueueRegionInfo retrieves all infos of one region on a service
func (c *Client) DBaasQueueRegionInfo(serviceName, regionID string) (*DBaasQueueRegion, error) {
	region := &DBaasQueueRegion{}
	err := c.Get(queryEscape("/dbaas/queue/%s/region/%s", serviceName, regionID), region)
	return region, err
}

// DBaasQueueTopicList list all topics on a service
func (c *Client) DBaasQueueTopicList(serviceName string, withDetails bool) ([]DBaasQueueTopic, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/topic", serviceName), &ids); err != nil {
		return nil, err
	}

	topics := []DBaasQueueTopic{}
	for _, id := range ids {
		topics = append(topics, DBaasQueueTopic{ID: id})
	}

	if !withDetails {
		return topics, nil
	}

	topicsChan, errChan := make(chan DBaasQueueTopic), make(chan error)
	for _, topic := range topics {
		go func(topic DBaasQueueTopic) {
			d, err := c.DBaasQueueTopicInfo(serviceName, topic.ID)
			if err != nil {
				errChan <- err
				return
			}
			topicsChan <- *d
		}(topic)
	}

	topicsComplete := []DBaasQueueTopic{}
	for i := 0; i < len(topics); i++ {
		select {
		case topics := <-topicsChan:
			topicsComplete = append(topicsComplete, topics)
		case err := <-errChan:
			return nil, err
		}
	}

	return topicsComplete, nil
}

// DBaasQueueTopicInfo retrieves all infos of one topic on a service
func (c *Client) DBaasQueueTopicInfo(serviceName, topicID string) (*DBaasQueueTopic, error) {
	topic := &DBaasQueueTopic{}
	err := c.Get(queryEscape("/dbaas/queue/%s/topic/%s", serviceName, topicID), topic)
	return topic, err
}

// DBaasQueueUserList list all users on a service
func (c *Client) DBaasQueueUserList(serviceName string, withDetails bool) ([]DBaasQueueUser, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/user", serviceName), &ids); err != nil {
		return nil, err
	}

	users := []DBaasQueueUser{}
	for _, id := range ids {
		users = append(users, DBaasQueueUser{ID: id})
	}

	if !withDetails {
		return users, nil
	}

	usersChan, errChan := make(chan DBaasQueueUser), make(chan error)
	for _, user := range users {
		go func(user DBaasQueueUser) {
			d, err := c.DBaasQueueUserInfo(serviceName, user.ID)
			if err != nil {
				errChan <- err
				return
			}
			usersChan <- *d
		}(user)
	}

	usersComplete := []DBaasQueueUser{}
	for i := 0; i < len(users); i++ {
		select {
		case users := <-usersChan:
			usersComplete = append(usersComplete, users)
		case err := <-errChan:
			return nil, err
		}
	}

	return usersComplete, nil
}

// DBaasQueueUserInfo retrieve all infos of one user of your apps
func (c *Client) DBaasQueueUserInfo(serviceName, userID string) (*DBaasQueueUser, error) {
	user := &DBaasQueueUser{}
	err := c.Get(queryEscape("/dbaas/queue/%s/user/%s", serviceName, userID), user)
	return user, err
}

// DBaasQueueUserChangePassword reset user password
func (c *Client) DBaasQueueUserChangePassword(serviceName, userID string) (*DBaasQueueUser, error) {
	user := &DBaasQueueUser{}
	err := c.Post(queryEscape("/dbaas/queue/%s/user/%s/changePassword", serviceName, userID), nil, user)
	return user, err
}

// DBaasQueueMetricsAccount retrieve all infos of one of your apps
func (c *Client) DBaasQueueMetricsAccount(serviceName string) (*DBaasQueueMetricsAccount, error) {
	user := &DBaasQueueMetricsAccount{}
	err := c.Get(queryEscape("/dbaas/queue/%s/metrics/account", serviceName), user)
	return user, err
}
