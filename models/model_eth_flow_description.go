package models

// Identifies an Ethernet flow
type EthFlowDescription struct {
	DestMacAddr string `json:"destMacAddr,omitempty" yaml:"destMacAddr" bson:"destMacAddr" mapstructure:"DestMacAddr"`
	EthType     string `json:"ethType" yaml:"ethType" bson:"ethType" mapstructure:"EthType"`
	// Defines a packet filter of an IP flow.
	FDesc         string        `json:"fDesc,omitempty" yaml:"fDesc" bson:"fDesc" mapstructure:"FDesc"`
	FDir          FlowDirection `json:"fDir,omitempty" yaml:"fDir" bson:"fDir" mapstructure:"FDir"`
	SourceMacAddr string        `json:"sourceMacAddr,omitempty" yaml:"sourceMacAddr" bson:"sourceMacAddr" mapstructure:"SourceMacAddr"`
	VlanTags      []string      `json:"vlanTags,omitempty" yaml:"vlanTags" bson:"vlanTags" mapstructure:"VlanTags"`
}
