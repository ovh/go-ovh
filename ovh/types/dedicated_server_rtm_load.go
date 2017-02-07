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

// DedicatedServerRtmLoad A structure describing informations about server load
type DedicatedServerRtmLoad struct {
	CPU *DedicatedServerRtmLoadCPU `json:"cpu,omitempty"`

	// Loadavg1 Load average in the last 1 minute
	Loadavg1 float64 `json:"loadavg1,omitempty"`

	// Loadavg5 Load average in the last 5 minutes
	Loadavg5 float64 `json:"loadavg5,omitempty"`

	// Loadavg15 Load average in the last 15 minutes
	Loadavg15 float64 `json:"loadavg15,omitempty"`

	Memory *DedicatedServerRtmLoadMemory `json:"memory,omitempty"`

	// ProcessCount Number of processes using or waiting for CPU time
	ProcessCount int64 `json:"processCount,omitempty"`

	// ProcessRunning Number of process running
	ProcessRunning int64 `json:"processRunning,omitempty"`

	Swap *DedicatedServerRtmLoadSwap `json:"swap,omitempty"`

	// Uptime Server uptime
	Uptime int64 `json:"uptime,omitempty"`
}
