package models

type Ncgi struct {
	PlmnId   *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	NrCellId string  `json:"nrCellId" yaml:"nrCellId" bson:"nrCellId"`
}
