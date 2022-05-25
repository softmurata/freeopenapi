package models

type RequestedUsageData struct {
	// An array of usage monitoring data id references to the usage monitoring data instances for which the PCF is requesting a usage report. This attribute shall only be provided when allUmIds is not set to true.
	RefUmIds []string `json:"refUmIds,omitempty" yaml:"refUmIds" bson:"refUmIds" mapstructure:"RefUmIds"`
	// Th ooleanean indicates whether requested usage data applies to all usage monitoring data instances. When it's not included, it means requested usage data shall only apply to the usage monitoring data instances referenced by the refUmIds attribute.
	AllUmIds bool `json:"allUmIds,omitempty" yaml:"allUmIds" bson:"allUmIds" mapstructure:"AllUmIds"`
}
