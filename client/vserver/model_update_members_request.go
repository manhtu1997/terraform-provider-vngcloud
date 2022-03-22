/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type UpdateMembersRequest struct {
	// List of members of the pool.
	Members []CreateMemberRequest `json:"members,omitempty"`
	// Id of project
	ProjectId string `json:"projectId,omitempty"`
	// Id of user
	UserId int32 `json:"userId,omitempty"`
}
