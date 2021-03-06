/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// NodePoolAutoScaling Node pool auto scaling settings.
type NodePoolAutoScaling struct {
	// Enable node pool autoscaling.
	Enabled bool `json:"enabled,omitempty"`
	// Minimum node pool size.
	MinSize int32 `json:"minSize"`
	// Maximum node pool size.
	MaxSize int32 `json:"maxSize"`
}
