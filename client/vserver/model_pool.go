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

type Pool struct {
	AdminStateUp bool `json:"adminStateUp,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Description string `json:"description,omitempty"`
	LoadBalanceMethod string `json:"loadBalanceMethod,omitempty"`
	LoadBalancerId string `json:"loadBalancerId,omitempty"`
	Name string `json:"name,omitempty"`
	OperatingStatus string `json:"operatingStatus,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	ProvisioningStatus string `json:"provisioningStatus,omitempty"`
	SessionPersistence int32 `json:"sessionPersistence,omitempty"`
	Status string `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
