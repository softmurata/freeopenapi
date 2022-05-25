package models

type QosData struct {
	// Univocally identifies the QoS control policy data within a PDU session.
	QosId   string `json:"qosId" yaml:"qosId" bson:"qosId" mapstructure:"QosId"`
	Var5qi  int32  `json:"5qi,omitempty" yaml:"5qi" bson:"5qi" mapstructure:"Var5qi"`
	MaxbrUl string `json:"maxbrUl,omitempty" yaml:"maxbrUl" bson:"maxbrUl" mapstructure:"MaxbrUl"`
	MaxbrDl string `json:"maxbrDl,omitempty" yaml:"maxbrDl" bson:"maxbrDl" mapstructure:"MaxbrDl"`
	GbrUl   string `json:"gbrUl,omitempty" yaml:"gbrUl" bson:"gbrUl" mapstructure:"GbrUl"`
	GbrDl   string `json:"gbrDl,omitempty" yaml:"gbrDl" bson:"gbrDl" mapstructure:"GbrDl"`
	Arp     *Arp   `json:"arp,omitempty" yaml:"arp" bson:"arp" mapstructure:"Arp"`
	// Indicates whether notifications are requested from 3GPP NG-RAN when the GFBR can no longer (or again) be guaranteed for a QoS Flow during the lifetime of the QoS Flow.
	Qnc             bool  `json:"qnc,omitempty" yaml:"qnc" bson:"qnc" mapstructure:"Qnc"`
	PriorityLevel   int32 `json:"priorityLevel,omitempty" yaml:"priorityLevel" bson:"priorityLevel" mapstructure:"PriorityLevel"`
	AverWindow      int32 `json:"averWindow,omitempty" yaml:"averWindow" bson:"averWindow" mapstructure:"AverWindow"`
	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty" yaml:"maxDataBurstVol" bson:"maxDataBurstVol" mapstructure:"MaxDataBurstVol"`
	// Indicates whether the QoS information is reflective for the corresponding service data flow.
	ReflectiveQos bool `json:"reflectiveQos,omitempty" yaml:"reflectiveQos" bson:"reflectiveQos" mapstructure:"ReflectiveQos"`
	// Indicates, by containing the same value, what PCC rules may share resource in downlink direction.
	SharingKeyDl string `json:"sharingKeyDl,omitempty" yaml:"sharingKeyDl" bson:"sharingKeyDl" mapstructure:"SharingKeyDl"`
	// Indicates, by containing the same value, what PCC rules may share resource in uplink direction.
	SharingKeyUl        string `json:"sharingKeyUl,omitempty" yaml:"sharingKeyUl" bson:"sharingKeyUl" mapstructure:"SharingKeyUl"`
	MaxPacketLossRateDl int32  `json:"maxPacketLossRateDl,omitempty" yaml:"maxPacketLossRateDl" bson:"maxPacketLossRateDl" mapstructure:"MaxPacketLossRateDl"`
	MaxPacketLossRateUl int32  `json:"maxPacketLossRateUl,omitempty" yaml:"maxPacketLossRateUl" bson:"maxPacketLossRateUl" mapstructure:"MaxPacketLossRateUl"`
	// Indicates that the dynamic PCC rule shall always have its binding with the QoS Flow associated with the default QoS rule
	DefQosFlowIndication bool `json:"defQosFlowIndication,omitempty" yaml:"defQosFlowIndication" bson:"defQosFlowIndication" mapstructure:"DefQosFlowIndication"`
}
