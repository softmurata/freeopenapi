/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type LadnInfo struct {
	Ladn     string        `json:"ladn"`
	Presence PresenceState `json:"presence,omitempty"`
}
