package models

type ServiceParameterData struct {
	AfServiceId     string  `json:"afServiceId,omitempty" yaml:"afServiceId" bson:"afServiceId" mapstructure:"AfServiceId"`
	AppId           string  `json:"appId,omitempty" yaml:"appId" bson:"appId" mapstructure:"AppId"`
	Dnn             string  `json:"dnn,omitempty" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
	Snssai          *Snssai `json:"snssai,omitempty" yaml:"snssai" bson:"snssai" mapstructure:"Snssai"`
	ExternalGroupId string  `json:"externalGroupId,omitempty" yaml:"externalGroupId" bson:"externalGroupId" mapstructure:"ExternalGroupId"`
	AnyUeInd        bool    `json:"anyUeInd,omitempty" yaml:"anyUeInd" bson:"anyUeInd" mapstructure:"AnyUeInd"`
	Gpsi            string  `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi" mapstructure:"Gpsi"`
	UeIpv4          string  `json:"ueIpv4,omitempty" yaml:"ueIpv4" bson:"ueIpv4" mapstructure:"UeIpv4"`
	UeIpv6          string  `json:"ueIpv6,omitempty" yaml:"ueIpv6" bson:"ueIpv6" mapstructure:"UeIpv6"`
	UeMac           string  `json:"ueMac,omitempty" yaml:"ueMac" bson:"ueMac" mapstructure:"UeMac"`
	ParamOverPc5    string  `json:"paramOverPc5,omitempty" yaml:"paramOverPc5" bson:"paramOverPc5" mapstructure:"ParamOverPc5"`
	ParamOverUu     string  `json:"paramOverUu,omitempty" yaml:"paramOverUu" bson:"paramOverUu" mapstructure:"ParamOverUu"`
	SuppFeat        string  `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat" mapstructure:"SuppFeat"`
}
