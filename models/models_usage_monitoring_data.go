package models

import (
	"time"
)

type UsageMonitoringData struct {
	// Univocally identifies the usage monitoring policy data within a PDU session.
	UmId string `json:"umId" yaml:"umId" bson:"umId" mapstructure:"UmId"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThreshold int64 `json:"volumeThreshold,omitempty" yaml:"volumeThreshold" bson:"volumeThreshold" mapstructure:"VolumeThreshold"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThresholdUplink int64 `json:"volumeThresholdUplink,omitempty" yaml:"volumeThresholdUplink" bson:"volumeThresholdUplink" mapstructure:"VolumeThresholdUplink"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThresholdDownlink int64      `json:"volumeThresholdDownlink,omitempty" yaml:"volumeThresholdDownlink" bson:"volumeThresholdDownlink" mapstructure:"VolumeThresholdDownlink"`
	TimeThreshold           int32      `json:"timeThreshold,omitempty" yaml:"timeThreshold" bson:"timeThreshold" mapstructure:"TimeThreshold"`
	MonitoringTime          *time.Time `json:"monitoringTime,omitempty" yaml:"monitoringTime" bson:"monitoringTime" mapstructure:"MonitoringTime"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThreshold int64 `json:"nextVolThreshold,omitempty" yaml:"nextVolThreshold" bson:"nextVolThreshold" mapstructure:"NextVolThreshold"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThresholdUplink int64 `json:"nextVolThresholdUplink,omitempty" yaml:"nextVolThresholdUplink" bson:"nextVolThresholdUplink" mapstructure:"NextVolThresholdUplink"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThresholdDownlink int64 `json:"nextVolThresholdDownlink,omitempty" yaml:"nextVolThresholdDownlink" bson:"nextVolThresholdDownlink" mapstructure:"NextVolThresholdDownlink"`
	NextTimeThreshold        int32 `json:"nextTimeThreshold,omitempty" yaml:"nextTimeThreshold" bson:"nextTimeThreshold" mapstructure:"NextTimeThreshold"`
	InactivityTime           int32 `json:"inactivityTime,omitempty" yaml:"inactivityTime" bson:"inactivityTime" mapstructure:"InactivityTime"`
	// Contains the PCC rule identifier(s) which corresponding service data flow(s) shall be excluded from PDU Session usage monitoring. It is only included in the UsageMonitoringData instance for session level usage monitoring.
	ExUsagePccRuleIds []string `json:"exUsagePccRuleIds,omitempty" yaml:"exUsagePccRuleIds" bson:"exUsagePccRuleIds" mapstructure:"ExUsagePccRuleIds"`
}
