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

// A bundle containing a set of policies, whitelists, and rules for mapping them to specific images
type PolicyBundle struct {
	// Id of the bundle
	Id string `json:"id"`
	// Human readable name for the bundle
	Name string `json:"name,omitempty"`
	// Description of the bundle, human readable
	Comment string `json:"comment,omitempty"`
	// Version id for this bundle format
	Version string `json:"version"`
	// Whitelists which define which policy matches to disregard explicitly in the final policy decision
	Whitelists []Whitelist `json:"whitelists,omitempty"`
	// Policies which define the go/stop/warn status of an image using rule matches on image properties
	Policies []Policy `json:"policies"`
	// Mapping rules for defining which policy and whitelist(s) to apply to an image based on a match of the image tag or id. Evaluated in order.
	Mappings []MappingRule `json:"mappings"`
	// List of mapping rules that define which images should always be passed (unless also on the blacklist), regardless of policy result.
	WhitelistedImages []ImageSelectionRule `json:"whitelistedImages,omitempty"`
	// List of mapping rules that define which images should always result in a STOP/FAIL policy result regardless of policy content or presence in whitelisted_images
	BlacklistedImages []ImageSelectionRule `json:"blacklistedImages,omitempty"`
}
