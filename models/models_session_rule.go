package models

type SessionRule struct {
	AuthSessAmbr *Ambr                 `json:"authSessAmbr,omitempty" yaml:"authSessAmbr" bson:"authSessAmbr" mapstructure:"AuthSessAmbr"`
	AuthDefQos   *AuthorizedDefaultQos `json:"authDefQos,omitempty" yaml:"authDefQos" bson:"authDefQos" mapstructure:"AuthDefQos"`
	// Univocally identifies the session rule within a PDU session.
	SessRuleId string `json:"sessRuleId" yaml:"sessRuleId" bson:"sessRuleId" mapstructure:"SessRuleId"`
	// A reference to UsageMonitoringData policy decision type. It is the umId described in subclause 5.6.2.12.
	RefUmData string `json:"refUmData,omitempty" yaml:"refUmData" bson:"refUmData" mapstructure:"RefUmData"`
	// A reference to the condition data. It is the condId described in subclause 5.6.2.9.
	RefCondData string `json:"refCondData,omitempty" yaml:"refCondData" bson:"refCondData" mapstructure:"RefCondData"`
}
