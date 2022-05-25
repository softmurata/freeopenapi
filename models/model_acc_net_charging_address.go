package models

// Describes the network entity within the access network performing charging
type AccNetChargingAddress struct {
	AnChargIpv4Addr string `json:"anChargIpv4Addr,omitempty" yaml:"anChargIpv4Addr" bson:"anChargIpv4Addr" mapstructure:"AnChargIpv4Addr"`
	AnChargIpv6Addr string `json:"anChargIpv6Addr,omitempty" yaml:"anChargIpv6Addr" bson:"anChargIpv6Addr" mapstructure:"AnChargIpv6Addr"`
}