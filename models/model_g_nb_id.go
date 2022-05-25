package models

type GNbId struct {
	BitLength int32  `json:"bitLength" yaml:"bitLength" bson:"bitLength"`
	GNBValue  string `json:"gNBValue" yaml:"gNBValue" bson:"gNBValue"`
}
