/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type ReorderPolicyRequest struct {
	// Policy's id to reorder
	PolicyId string `json:"policyId,omitempty"`
	// New position of the policy
	Position int64 `json:"position,omitempty"`
}
