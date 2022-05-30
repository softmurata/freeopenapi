package models

type QosMonitoringInformation struct {
	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams" yaml:"reqQosMonParams" bson:"reqQosMonParams" mapstructure:"ReqQosMonParams"`
	RepFreqs        []ReportingFrequency              `json:"repFreqs" yaml:"repFreqs" bson:"repFreqs" mapstructure:"RepFreqs"`
	RepThreshDl     int                               `json:"repThreshDl,omitempty" yaml:"repThreshDl" bson:"repThreshDl" mapstructure:"RepThreshDl"`
	RepThreshUl     int                               `json:"repThreshUl,omitempty" yaml:"repThreshUl" bson:"repThreshUl" mapstructure:"RepThreshUl"`
	RepThreshRp     int                               `json:"repThreshRp,omitempty" yaml:"repThreshRp" bson:"repThreshRp" mapstructure:"RepThreshRp"`
	RepPeriod       int                               `json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod" mapstructure:"RepPeriod"`
	WaitTime        int                               `json:"waitTime" yaml:"waitTime" bson:"waitTime" mapstructure:"WaitTime"`
}
