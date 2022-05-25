package models

type FlowInformation struct {
	// Defines a packet filter for an IP flow.Refer to subclause 5.4.2 of 3GPP TS 29.212 [23] for encoding.
	FlowDescription    string              `json:"flowDescription,omitempty" yaml:"flowDescription" bson:"flowDescription" mapstructure:"FlowDescription"`
	EthFlowDescription *EthFlowDescription `json:"ethFlowDescription,omitempty" yaml:"ethFlowDescription" bson:"ethFlowDescription" mapstructure:"EthFlowDescription"`
	// An identifier of packet filter.
	PackFiltId string `json:"packFiltId,omitempty" yaml:"packFiltId" bson:"packFiltId" mapstructure:"PackFiltId"`
	// The packet shall be sent to the UE.
	PacketFilterUsage bool `json:"packetFilterUsage,omitempty" yaml:"packetFilterUsage" bson:"packetFilterUsage" mapstructure:"PacketFilterUsage"`
	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass string `json:"tosTrafficClass,omitempty" yaml:"tosTrafficClass" bson:"tosTrafficClass" mapstructure:"TosTrafficClass"`
	// the security parameter index of the IPSec packet.
	Spi string `json:"spi,omitempty" yaml:"spi" bson:"spi" mapstructure:"Spi"`
	// the Ipv6 flow label header field.
	FlowLabel     string          `json:"flowLabel,omitempty" yaml:"flowLabel" bson:"flowLabel" mapstructure:"FlowLabel"`
	FlowDirection FlowDirectionRm `json:"flowDirection,omitempty" yaml:"flowDirection" bson:"flowDirection" mapstructure:"FlowDirection"`
}
