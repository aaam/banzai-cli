/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.29.0-dev.1
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type CreateUpdateOkePropertiesOke struct {
	Version string `json:"version"`
	NodePools map[string]NodePoolsOracle `json:"nodePools"`
	VcnId string `json:"vcnId,omitempty"`
	LbSubnetId1 string `json:"lbSubnetId1,omitempty"`
	LbSubnetId2 string `json:"lbSubnetId2,omitempty"`
}
