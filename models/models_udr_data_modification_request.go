package models

type UdrDataModificationNotifyRequest struct {
	ResUri string  `json:"resUri", yaml:"resUri", bson:"resUri", mapstructure:"ResUri"`
	ServParamData *ServParamDataRequest `json:"servParamData", yaml:"servParamData", bson:"servParamData", mapstructure:"ServParamData"`
}
