/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type Quota struct {
	Description string `json:"description,omitempty"`
	Limit int64 `json:"limit,omitempty"`
	ProjectUuid string `json:"projectUuid,omitempty"`
	QuotaName string `json:"quotaName,omitempty"`
	Type_ string `json:"type,omitempty"`
	Used string `json:"used,omitempty"`
}
