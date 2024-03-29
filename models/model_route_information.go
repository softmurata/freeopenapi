package models

type RouteInformation struct {
	Ipv4Addr   string `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr" mapstructure:"Ipv4Addr"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty" yaml:"ipv6Addr" bson:"ipv6Addr" mapstructure:"Ipv6Addr"`
	PortNumber int32  `json:"portNumber" yaml:"portNumber" bson:"portNumber" mapstructure:"PortNumber"`
}
