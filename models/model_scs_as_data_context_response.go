package models

type ScsAsDataContextResponse struct {
	Self                    string                   `json:"self,omitempty" yaml:"self" bson:"self" mapstructure:"Self"`
	FlowInfos               []FlowInfo               `json:"flowInfos,omitempty" yaml:"flowInfos" bson:"flowInfos" mapstructure:"FlowInfos"`
	NotificationDestination string                   `json:"notificationDestination" yaml:"notificationDestination" bson:"notificationDestination" mapstructure:"NotificationDestination"`
	EthFlowInfos            []EthFlowInfo            `json:"ethFlowInfos,omitempty" yaml:"ethFlowInfos" bson:"ethFlowInfos" mapstructure:"EthFlowInfos"`
	QosReference            string                   `json:"qosReference,omitempty" yaml:"qosReference" bson:"qosReference" mapstructure:"QosReference"`
	UeIpv4Addr              string                   `json:"ueIpv4Addr,omitempty" yaml:"ueIpv4Addr" bson:"ueIpv4Addr" mapstructure:"UeIpv4Addr"`
	UeIpv6Addr              string                   `json:"ueIpv6Addr,omitempty" yaml:"ueIpv6Addr" bson:"ueIpv6Addr" mapstructure:"UeIpv6Addr"`
	MacAddr                 string                   `json:"macAddr,omitempty" yaml:"macAddr" bson:"macAddr" mapstructure:"MacAddr"`
	QosMonInfo              QosMonitoringInformation `json:"qosMonInfo,omitempty" yaml:"qosMonInfo" bson:"qosMonInfo" mapstructure:"QosMonInfo"`
	SupportedFeatures       string                   `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
	Supi                    string                   `json:"supi,omitempty" yaml:"supi" bson:"supi" mapstructure:"Supi"` // for loadcore
	Dnn                     string                   `json:"dnn,omitempty" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`     // for free5gc
}
