package models

type AccNetChId struct {
	AccNetChaIdValue int32 `json:"accNetChaIdValue" yaml:"accNetChaIdValue" bson:"accNetChaIdValue" mapstructure:"AccNetChaIdValue"`
	// Contains the identifier of the PCC rule(s) associated to the provided Access Network Charging Identifier.
	RefPccRuleIds []string `json:"refPccRuleIds,omitempty" yaml:"refPccRuleIds" bson:"refPccRuleIds" mapstructure:"RefPccRuleIds"`
	// When it is included and set to true, indicates the Access Network Charging Identifier applies to the whole PDU Session
	SessionChScope bool `json:"sessionChScope,omitempty" yaml:"sessionChScope" bson:"sessionChScope" mapstructure:"SessionChScope"`
}
