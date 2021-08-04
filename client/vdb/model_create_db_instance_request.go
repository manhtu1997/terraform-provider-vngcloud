/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vdb

import (
	"time"
)

type CreateDbInstanceRequest struct {
	AutoRenewPeriod      int32             `json:"autoRenewPeriod,omitempty"`
	BackupAuto           bool              `json:"backupAuto,omitempty"`
	BackupDuration       int32             `json:"backupDuration,omitempty"`
	BackupTime           string            `json:"backupTime,omitempty"`
	CartItemId           int32             `json:"cartItemId,omitempty"`
	CartItemState        int32             `json:"cartItemState,omitempty"`
	ConfigId             string            `json:"configId"`
	Cost                 float64           `json:"cost,omitempty"`
	Databases            []DatabaseRequest `json:"databases,omitempty"`
	DatastoreType        string            `json:"datastoreType,omitempty"`
	DatastoreVersion     string            `json:"datastoreVersion,omitempty"`
	EnableAutoRenew      bool              `json:"enableAutoRenew"`
	EndTime              *time.Time        `json:"endTime,omitempty"`
	EngineGroup          int32             `json:"engineGroup,omitempty"`
	Extra                interface{}       `json:"extra,omitempty"`
	FlavorId             string            `json:"flavorId,omitempty"`
	Id                   string            `json:"id,omitempty"`
	InvoiceId            int32             `json:"invoiceId,omitempty"`
	Name                 string            `json:"name,omitempty"`
	NetIds               []string          `json:"netIds,omitempty"`
	PackageId            string            `json:"packageId,omitempty"`
	Period               int32             `json:"period,omitempty"`
	Poc                  bool              `json:"poc,omitempty"`
	PriceKey             string            `json:"priceKey,omitempty"`
	ProjectId            string            `json:"projectId,omitempty"`
	PublicAccess         bool              `json:"publicAccess,omitempty"`
	Ram                  int32             `json:"ram,omitempty"`
	RedisPassword        string            `json:"redisPassword,omitempty"`
	RedisPasswordEnabled bool              `json:"redisPasswordEnabled"`
	ReplicaSourceId      string            `json:"replicaSourceId,omitempty"`
	StartTime            time.Time         `json:"startTime,omitempty"`
	UseTrial             bool              `json:"useTrial,omitempty"`
	User                 *UserRequest      `json:"user,omitempty"`
	Vcpus                int32             `json:"vcpus,omitempty"`
	VolumeSize           int32             `json:"volumeSize,omitempty"`
	VolumeType           string            `json:"volumeType,omitempty"`
	VolumeTypeZoneId     string            `json:"volumeTypeZoneId,omitempty"`
	ZoneId               int32             `json:"zoneId,omitempty"`
}
