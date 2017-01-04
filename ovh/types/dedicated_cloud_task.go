/* 
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

import (
	"time"
)

// Operation on a Dedicated Cloud component
type DedicatedCloudTask struct {

	// Creator of the task
	CreatedBy string `json:"createdBy,omitempty"`

	// Origin of the task
	CreatedFrom string `json:"createdFrom,omitempty"`

	// datacenterId of the associated dedicatedCloud.Datacenter object
	DatacenterId int64 `json:"datacenterId,omitempty"`

	// Current progress description
	Description string `json:"description,omitempty"`

	// Task end date
	EndDate time.Time `json:"endDate,omitempty"`

	// Task execution date
	ExecutionDate time.Time `json:"executionDate,omitempty"`

	// filerId of the associated dedicatedCloud.Filer object
	FilerId int64 `json:"filerId,omitempty"`

	// hostId of the associated dedicatedCloud.Host object
	HostId int64 `json:"hostId,omitempty"`

	// Task last modification date
	LastModificationDate time.Time `json:"lastModificationDate,omitempty"`

	// Maintenance task min allowed execution date
	MaintenanceDateFrom time.Time `json:"maintenanceDateFrom,omitempty"`

	// Maintenance task max allowed execution date
	MaintenanceDateTo time.Time `json:"maintenanceDateTo,omitempty"`

	// Task name
	Name string `json:"name,omitempty"`

	// network of the associated dedicatedCloud.Ip object
	Network string `json:"network,omitempty"`

	// networkAccessId of the associated dedicatedCloud.AllowedNetwork object
	NetworkAccessId int64 `json:"networkAccessId,omitempty"`

	// orderId of the associated billing.Order object
	OrderId int64 `json:"orderId,omitempty"`

	// taskId of the parent dedicatedCloud.Task object
	ParentTaskId int64 `json:"parentTaskId,omitempty"`

	// Current progress
	Progress int64 `json:"progress,omitempty"`

	// Current Task state
	State string `json:"state,omitempty"`

	// Task id
	TaskId int64 `json:"taskId,omitempty"`

	// Task type
	Type_ string `json:"type,omitempty"`

	// userId of the associated dedicatedCloud.User object
	UserId int64 `json:"userId,omitempty"`

	// vlanId of the parent dedicatedCloud.Vlan object
	VlanId int64 `json:"vlanId,omitempty"`
}
