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

type NodePoolsAmazon struct {
	InstanceType string `json:"instanceType"`
	SpotPrice string `json:"spotPrice"`
	Autoscaling bool `json:"autoscaling,omitempty"`
	Count int32 `json:"count,omitempty"`
	MinCount int32 `json:"minCount"`
	MaxCount int32 `json:"maxCount"`
	Labels map[string]string `json:"labels,omitempty"`
	Image string `json:"image,omitempty"`
}
