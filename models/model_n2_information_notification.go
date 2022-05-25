/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type N2InformationNotification struct {
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
	N2InfoContainer        *N2InfoContainer   `json:"n2InfoContainer,omitempty"`
	ToReleaseSessionList   []int32            `json:"toReleaseSessionList,omitempty"`
	LcsCorrelationId       string             `json:"lcsCorrelationId,omitempty"`
	NotifyReason           N2InfoNotifyReason `json:"notifyReason,omitempty"`
}