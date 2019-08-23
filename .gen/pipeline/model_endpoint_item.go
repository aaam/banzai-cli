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

type EndpointItem struct {
	Name string `json:"name,omitempty"`
	Host string `json:"host,omitempty"`
	Urls []UrlItem `json:"urls,omitempty"`
}
