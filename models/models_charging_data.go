package models

type ChargingData struct {
	// Univocally identifies the charging control policy data within a PDU session.
	ChgId          string         `json:"chgId" yaml:"chgId" bson:"chgId" mapstructure:"ChgId"`
	MeteringMethod MeteringMethod `json:"meteringMethod,omitempty" yaml:"meteringMethod" bson:"meteringMethod" mapstructure:"MeteringMethod"`
	// Indicates the offline charging is applicable to the PDU session or PCC rule.
	Offline bool `json:"offline,omitempty" yaml:"offline" bson:"offline" mapstructure:"Offline"`
	// Indicates the online charging is applicable to the PDU session or PCC rule.
	Online bool `json:"online,omitempty" yaml:"online" bson:"online" mapstructure:"Online"`
	// Indicates whether the service data flow is allowed to start while the SMF is waiting for the response to the credit request.
	SdfHandl       bool           `json:"sdfHandl,omitempty" yaml:"sdfHandl" bson:"sdfHandl" mapstructure:"SdfHandl"`
	RatingGroup    int32          `json:"ratingGroup,omitempty" yaml:"ratingGroup" bson:"ratingGroup" mapstructure:"RatingGroup"`
	ReportingLevel ReportingLevel `json:"reportingLevel,omitempty" yaml:"reportingLevel" bson:"reportingLevel" mapstructure:"ReportingLevel"`
	ServiceId      int32          `json:"serviceId,omitempty" yaml:"serviceId" bson:"serviceId" mapstructure:"ServiceId"`
	// Indicates the sponsor identity.
	SponsorId string `json:"sponsorId,omitempty" yaml:"sponsorId" bson:"sponsorId" mapstructure:"SponsorId"`
	// Indicates the application service provider identity.
	AppSvcProvId         string `json:"appSvcProvId,omitempty" yaml:"appSvcProvId" bson:"appSvcProvId" mapstructure:"AppSvcProvId"`
	AfChargingIdentifier int32  `json:"afChargingIdentifier,omitempty" yaml:"afChargingIdentifier" bson:"afChargingIdentifier" mapstructure:"AfChargingIdentifier"`
}
