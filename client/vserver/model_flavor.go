/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type Flavor struct {
	Bandwidth     int32  `json:"bandwidth,omitempty"`
	BandwidthUnit string `json:"bandwidthUnit,omitempty"`
	Cpu           int32  `json:"cpu,omitempty"`
	FlavorId      string `json:"flavorId,omitempty"`
	Gpu           int32  `json:"gpu,omitempty"`
	Memory        int32  `json:"memory,omitempty"`
	Name          string `json:"name,omitempty"`
	ZoneId        string `json:"zoneId,omitempty"`
}
