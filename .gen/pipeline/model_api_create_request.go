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

type ApiCreateRequest struct {
	Members []int32 `json:"members,omitempty"`
	Name string `json:"name,omitempty"`
}
