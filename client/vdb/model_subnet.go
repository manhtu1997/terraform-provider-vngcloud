/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vdb

type Subnet struct {
	Name      string `json:"name,omitempty"`
	ProjectId int64  `json:"projectId,omitempty"`
	Subnet    string `json:"subnet,omitempty"`
	SubnetId  int64  `json:"subnetId,omitempty"`
	VpcId     int64  `json:"vpcId,omitempty"`
}
