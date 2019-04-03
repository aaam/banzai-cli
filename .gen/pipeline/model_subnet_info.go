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

type SubnetInfo struct {
	// The IPv4 CIDR blocks assigned to the subnet
	Cidrs []string `json:"cidrs"`
	// Identifier of the subnetwork
	Id string `json:"id"`
	// The location of the subnetwork.
	Location string `json:"location,omitempty"`
	// Name of the subnetwork
	Name string `json:"name,omitempty"`
}
