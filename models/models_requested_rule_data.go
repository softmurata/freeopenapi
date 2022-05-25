package models

type RequestedRuleData struct {
	// An array of PCC rule id references to the PCC rules associated with the control data.
	RefPccRuleIds []string `json:"refPccRuleIds" yaml:"refPccRuleIds" bson:"refPccRuleIds" mapstructure:"RefPccRuleIds"`
	// Array of requested rule data type elements indicating what type of rule data is requested for the corresponding referenced PCC rules.
	ReqData []RequestedRuleDataType `json:"reqData" yaml:"reqData" bson:"reqData" mapstructure:"ReqData"`
}
