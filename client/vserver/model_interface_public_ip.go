/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type InterfacePublicIp struct {
	BackendStatus string `json:"backendStatus,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	GatewayIp string `json:"gatewayIp,omitempty"`
	Id int32 `json:"id,omitempty"`
	Ip string `json:"ip,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
	Status string `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	VpcBackendId int32 `json:"vpcBackendId,omitempty"`
}
