package models

type Ecgi struct {
	PlmnId      *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	EutraCellId string  `json:"eutraCellId" yaml:"eutraCellId" bson:"eutraCellId"`
}
