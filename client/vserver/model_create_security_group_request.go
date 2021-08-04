/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateSecurityGroupRequest struct {
	// description for secgroup.
	Description string       `json:"description,omitempty"`
	Extra       *interface{} `json:"extra,omitempty"`
	// Name of the Secgroup
	Name string `json:"name"`
}
