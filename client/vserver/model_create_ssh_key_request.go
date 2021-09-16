/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateSshKeyRequest struct {
	Name string `json:"name,omitempty"`
	// Id of project
	ProjectId string `json:"projectId,omitempty"`
	// Id of user
	UserId int32 `json:"userId,omitempty"`
}
