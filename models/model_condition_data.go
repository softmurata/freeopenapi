package models

import (
	"time"
)

type ConditionData struct {
	// Uniquely identifies the condition data within a PDU session.
	CondId           string     `json:"condId" yaml:"condId" bson:"condId" mapstructure:"CondId"`
	ActivationTime   *time.Time `json:"activationTime,omitempty" yaml:"activationTime" bson:"activationTime" mapstructure:"ActivationTime"`
	DeactivationTime *time.Time `json:"deactivationTime,omitempty" yaml:"deactivationTime" bson:"deactivationTime" mapstructure:"DeactivationTime"`
}