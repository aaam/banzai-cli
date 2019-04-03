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

type CreateAckPropertiesAcsk struct {
	RegionId string `json:"regionId"`
	ZoneId string `json:"zoneId"`
	VswitchId string `json:"vswitchId,omitempty"`
	NodePools map[string]NodePoolsAck `json:"nodePools"`
}
