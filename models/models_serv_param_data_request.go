package models


type ServParamDataRequest struct {
	AnyUeInd bool `json:"anyUeInd", yaml:"anyUeInd", bson:"anyUeInd", mapstructure:"AnyUeInd"`
	AppId string `json:"appId", yaml:"appId", bson:"appId", mapstructure:"AppId"`
	Dnn string `json:"dnn", yaml:"dnn", bson:"dnn", mapstructure:"Dnn"`
	ParamOverPc5 string `json:"paramOverPc5", yaml:"paramOverPc5", bson:"paramOverPc5", mapstructure:"ParamOverPc5"`
	ParamOverUu string `json:"paramOverUu", yaml:"paramOverUu", bson:"paramOverUu", mapstructure:"ParamOverUu"`
	Snssai *Snssai `json:"snssai", yaml:"snssai", bson:"snssai", mapstructure:"Snssai"`
	Supi string `json:"supi", yaml:"supi", bson:"supi", mapstructure:"Supi"`
	SuppFeat string `json:"suppFeat", yaml:"suppFeat", bson:"suppFeat", mapstructure:"SuppFeat"`
	UeIpv4 string `json:"ueIpv4", yaml:"ueIpv4", bson:"ueIpv4", mapstructure:"UeIpv4"`

}
