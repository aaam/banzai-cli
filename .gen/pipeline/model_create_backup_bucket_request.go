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

type CreateBackupBucketRequest struct {
	Cloud string `json:"cloud"`
	BucketName string `json:"bucketName"`
	SecretId string `json:"secretId"`
}
