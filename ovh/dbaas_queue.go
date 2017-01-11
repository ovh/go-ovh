package ovh

import (
	"fmt"

	"github.com/runabove/go-sdk/ovh/types"
)

// DBaasQueueAppList list all your app
func (c *Client) DBaasQueueAppList(withDetails bool) ([]types.DBaasQueueApp, error) {
	var ids []string
	if err := c.Get("/dbaas/queue", &ids); err != nil {
		return nil, err
	}

	apps := []types.DBaasQueueApp{}
	for _, name := range ids {
		apps = append(apps, types.DBaasQueueApp{ID: name})
	}

	if !withDetails {
		return apps, nil
	}

	appsChan, errChan := make(chan types.DBaasQueueApp), make(chan error)
	for _, app := range apps {
		go func(app types.DBaasQueueApp) {
			d, err := c.DBaasQueueAppInfo(app.ID)
			if err != nil {
				errChan <- err
				return
			}
			appsChan <- *d
		}(app)
	}

	appsComplete := []types.DBaasQueueApp{}
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
func (c *Client) DBaasQueueAppInfo(serviceName string) (*types.DBaasQueueApp, error) {
	app := &types.DBaasQueueApp{}
	err := c.Get(queryEscape("/dbaas/queue/%s", serviceName), app)
	return app, err
}

// DBaasQueueAppServiceInfo retrieve all infos of one of your apps
func (c *Client) DBaasQueueAppServiceInfo(serviceName string) (*types.ServicesService, error) {
	app := &types.ServicesService{}
	err := c.Get(queryEscape("/dbaas/queue/%s/serviceInfos", serviceName), app)
	return app, err
}

// DBaasQueueAppInfoByName retrieve all infos of one of your apps
func (c *Client) DBaasQueueAppInfoByName(name string) (*types.DBaasQueueApp, error) {
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
func (c *Client) DBaasQueueKeyList(serviceName string, withDetails bool) ([]types.DBaasQueueKey, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/key", serviceName), &ids); err != nil {
		return nil, err
	}

	keys := []types.DBaasQueueKey{}
	for _, name := range ids {
		keys = append(keys, types.DBaasQueueKey{ID: name})
	}

	if !withDetails {
		return keys, nil
	}

	keysChan, errChan := make(chan types.DBaasQueueKey), make(chan error)
	for _, key := range keys {
		go func(key types.DBaasQueueKey) {
			d, err := c.DBaasQueueKeyInfo(serviceName, key.ID)
			if err != nil {
				errChan <- err
				return
			}
			keysChan <- *d
		}(key)
	}

	keysComplete := []types.DBaasQueueKey{}
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
func (c *Client) DBaasQueueKeyInfo(serviceName, keyID string) (*types.DBaasQueueKey, error) {
	key := &types.DBaasQueueKey{}
	err := c.Get(queryEscape("/dbaas/queue/%s/key/%s", serviceName, keyID), key)
	return key, err
}

// DBaasQueueRoleList list all roles on a service
func (c *Client) DBaasQueueRoleList(serviceName string, withDetails bool) ([]types.DBaasQueueRole, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/role", serviceName), &ids); err != nil {
		return nil, err
	}

	roles := []types.DBaasQueueRole{}
	for _, name := range ids {
		roles = append(roles, types.DBaasQueueRole{Name: name})
	}

	if !withDetails {
		return roles, nil
	}

	rolesChan, errChan := make(chan types.DBaasQueueRole), make(chan error)
	for _, role := range roles {
		go func(role types.DBaasQueueRole) {
			d, err := c.DBaasQueueRoleInfo(serviceName, role.Name)
			if err != nil {
				errChan <- err
				return
			}
			rolesChan <- *d
		}(role)
	}

	rolesComplete := []types.DBaasQueueRole{}
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
func (c *Client) DBaasQueueRoleInfo(serviceName, roleID string) (*types.DBaasQueueRole, error) {
	role := &types.DBaasQueueRole{}
	err := c.Get(queryEscape("/dbaas/queue/%s/role/%s", serviceName, roleID), role)
	return role, err
}

// DBaasQueueRegionList list all region on a service
func (c *Client) DBaasQueueRegionList(serviceName string, withDetails bool) ([]types.DBaasQueueRegion, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/region", serviceName), &ids); err != nil {
		return nil, err
	}

	regions := []types.DBaasQueueRegion{}
	for _, name := range ids {
		regions = append(regions, types.DBaasQueueRegion{Name: name})
	}

	if !withDetails {
		return regions, nil
	}

	regionsChan, errChan := make(chan types.DBaasQueueRegion), make(chan error)
	for _, region := range regions {
		go func(region types.DBaasQueueRegion) {
			d, err := c.DBaasQueueRegionInfo(serviceName, region.Name)
			if err != nil {
				errChan <- err
				return
			}
			regionsChan <- *d
		}(region)
	}

	regionsComplete := []types.DBaasQueueRegion{}
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
func (c *Client) DBaasQueueRegionInfo(serviceName, regionID string) (*types.DBaasQueueRegion, error) {
	region := &types.DBaasQueueRegion{}
	err := c.Get(queryEscape("/dbaas/queue/%s/region/%s", serviceName, regionID), region)
	return region, err
}

// DBaasQueueTopicList list all topics on a service
func (c *Client) DBaasQueueTopicList(serviceName string, withDetails bool) ([]types.DBaasQueueTopic, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/topic", serviceName), &ids); err != nil {
		return nil, err
	}

	topics := []types.DBaasQueueTopic{}
	for _, id := range ids {
		topics = append(topics, types.DBaasQueueTopic{ID: id})
	}

	if !withDetails {
		return topics, nil
	}

	topicsChan, errChan := make(chan types.DBaasQueueTopic), make(chan error)
	for _, topic := range topics {
		go func(topic types.DBaasQueueTopic) {
			d, err := c.DBaasQueueTopicInfo(serviceName, topic.ID)
			if err != nil {
				errChan <- err
				return
			}
			topicsChan <- *d
		}(topic)
	}

	topicsComplete := []types.DBaasQueueTopic{}
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
func (c *Client) DBaasQueueTopicInfo(serviceName, topicID string) (*types.DBaasQueueTopic, error) {
	topic := &types.DBaasQueueTopic{}
	err := c.Get(queryEscape("/dbaas/queue/%s/topic/%s", serviceName, topicID), topic)
	return topic, err
}

// DBaasQueueUserList list all users on a service
func (c *Client) DBaasQueueUserList(serviceName string, withDetails bool) ([]types.DBaasQueueUser, error) {
	var ids []string
	if err := c.Get(queryEscape("/dbaas/queue/%s/user", serviceName), &ids); err != nil {
		return nil, err
	}

	users := []types.DBaasQueueUser{}
	for _, id := range ids {
		users = append(users, types.DBaasQueueUser{ID: id})
	}

	if !withDetails {
		return users, nil
	}

	usersChan, errChan := make(chan types.DBaasQueueUser), make(chan error)
	for _, user := range users {
		go func(user types.DBaasQueueUser) {
			d, err := c.DBaasQueueUserInfo(serviceName, user.ID)
			if err != nil {
				errChan <- err
				return
			}
			usersChan <- *d
		}(user)
	}

	usersComplete := []types.DBaasQueueUser{}
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
func (c *Client) DBaasQueueUserInfo(serviceName, userID string) (*types.DBaasQueueUser, error) {
	user := &types.DBaasQueueUser{}
	err := c.Get(queryEscape("/dbaas/queue/%s/user/%s", serviceName, userID), user)
	return user, err
}

// DBaasQueueUserChangePassword reset user password
func (c *Client) DBaasQueueUserChangePassword(serviceName, userID string) (*types.DBaasQueueUser, error) {
	user := &types.DBaasQueueUser{}
	err := c.Post(queryEscape("/dbaas/queue/%s/user/%s/changePassword", serviceName, userID), nil, user)
	return user, err
}

// DBaasQueueMetricsAccount retrieve all infos of one of your apps
func (c *Client) DBaasQueueMetricsAccount(serviceName string) (*types.DBaasQueueMetricsAccount, error) {
	user := &types.DBaasQueueMetricsAccount{}
	err := c.Get(queryEscape("/dbaas/queue/%s/metrics/account", serviceName), user)
	return user, err
}
