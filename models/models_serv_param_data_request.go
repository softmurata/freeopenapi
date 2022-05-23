package models


type ServParamDataRequest struct {
	AnyUeInd bool `json:"anyUeInd,omitempty" yaml:"anyUeInd" bson:"anyUeInd" mapstructure:"AnyUeInd"`
	AppId string `json:"appId,omitempty" yaml:"appId" bson:"appId" mapstructure:"AppId"`
	Dnn string `json:"dnn,omitempty" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
	InterGroupId string `json:"interGroupId,omitempty" yaml:"interGroupId" bson:"interGroupId" mapstructure:"InterGroupId"`
	ParamOverPc5 string `json:"paramOverPc5" yaml:"paramOverPc5" bson:"paramOverPc5" mapstructure:"ParamOverPc5"`
	ParamOverUu string `json:"paramOverUu" yaml:"paramOverUu" bson:"paramOverUu" mapstructure:"ParamOverUu"`
	Snssai *Snssai `json:"snssai" yaml:"snssai" bson:"snssai" mapstructure:"Snssai"`
	Supi string `json:"supi,omitempty" yaml:"supi" bson:"supi" mapstructure:"Supi"`
	SuppFeat string `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat" mapstructure:"SuppFeat"`
	UeIpv4 string `json:"ueIpv4,omitempty" yaml:"ueIpv4" bson:"ueIpv4" mapstructure:"UeIpv4"`
	UeIpv6 string `json:"ueIpv6,omitempty" yaml:"ueIpv6" bson:"ueIpv6" mapStructure:"UeIpv6"`
	UeMac  string `json:"ueMac,omitempty" yaml:"ueMac" bson:"ueMac" mapStructure:"UeMac"`
}
