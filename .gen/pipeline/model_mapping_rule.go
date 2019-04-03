/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.15.4
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type MappingRule struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	WhitelistIds []string `json:"whitelistIds,omitempty"`
	// Optional single policy to evalute, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1.
	PolicyId string `json:"policyId,omitempty"`
	// List of policyIds to evaluate in order, to completion
	PolicyIds []string `json:"policyIds,omitempty"`
	Registry string `json:"registry"`
	Repository string `json:"repository"`
	Image ImageRef `json:"image"`
}
