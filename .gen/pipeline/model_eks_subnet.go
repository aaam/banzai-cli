/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.29.1
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type EksSubnet struct {
	// Id of existing subnet to use for creating the EKS cluster. If not provided new subnet will be created.
	SubnetId string `json:"subnetId,omitempty"`
	// The CIDR range for the subnet in case new Subnet is created.
	Cidr string `json:"cidr,omitempty"`
}
